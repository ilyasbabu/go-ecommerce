package services

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ilyasbabu/go-ecommerce/models"
	"gorm.io/gorm"
)

type ValidatedProductCreateData struct {
	Name        string
	Description string
	Price       float64
	Image       *multipart.FileHeader
	Path        string
}

func ValidateProductCreate(c *gin.Context) (ValidatedProductCreateData, error) {
	var data ValidatedProductCreateData
	name := c.PostForm("name")
	if name == "" {
		return data, errors.New("name missing")
	}
	data.Name = name
	description := c.PostForm("description")
	if description == "" {
		return data, errors.New("description missing")
	}
	data.Description = description
	price := c.PostForm("price")
	if price == "" {
		return data, errors.New("price missing")
	}
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return data, errors.New("invalid price")
	}
	data.Price = floatPrice

	file, imageHeader, err := c.Request.FormFile("image")
	if err != nil {
		return data, errors.New("image missing")
	}
	extension := filepath.Ext(imageHeader.Filename)
	newFileName := uuid.New().String() + extension
	dateTime := time.Now().Format("2006-01-02")
	path := "images/" + dateTime + "/" + newFileName
	data.Path = path
	if !IsImageFile(file) {
		return data, errors.New("image type invalid")
	}
	data.Image = imageHeader

	return data, nil
}

func CreateProduct(validatedData ValidatedProductCreateData, c *gin.Context) error {
	if err := c.SaveUploadedFile(validatedData.Image, "./public/"+validatedData.Path); err != nil {
		return errors.New("error saving image")
	}

	product := models.Products{Name: validatedData.Name, Description: validatedData.Description, Price: validatedData.Price}
	result := Db.Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return errors.New("Cannot Update! Product with same name(" + product.Name + ") already exists")
		} else {
			return result.Error
		}
	}

	product_image := models.ProductImages{ProductId: product.ID, Path: validatedData.Path, IsMain: true}
	result = Db.Create(&product_image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type ProductDetail struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createAt"`
	Slug        string    `json:"slug"`
	Stock       int       `json:"stock"`
	AvgRating   float64   `json:"avgRating"`
	RatingCount int       `json:"ratingCount"`
	Image       string    `json:"image"`
}

func GetProductByID(id string) (ProductDetail, error) {
	var product models.Products
	var data ProductDetail
	err := Db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("product not found")
	} else if err != nil {
		return data, err
	}
	var productImage models.ProductImages
	err = Db.Where("product_id= ? AND is_main=?", product.ID, true).First(&productImage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("image not found")
	}
	data = ProductDetail{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		Slug:        product.Slug,
		Stock:       product.Stock,
		AvgRating:   product.AvgRating,
		RatingCount: product.RatingCount,
		Image:       "media/" + productImage.Path,
	}
	return data, nil
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createAt"`
	Slug        string    `json:"slug"`
	Stock       int       `json:"stock"`
	AvgRating   float64   `json:"avgRating"`
	RatingCount int       `json:"ratingCount"`
	Image       string    `json:"image"`
}

func GetProducts() []ProductResponse {
	var products []models.Products

	Db.Find(&products)

	var responseData []ProductResponse

	for _, product := range products {
		var productImage models.ProductImages
		Db.Where("product_id= ? AND is_main=?", product.ID, true).First(&productImage)
		responseData = append(responseData, ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			Slug:        product.Slug,
			Stock:       product.Stock,
			AvgRating:   product.AvgRating,
			RatingCount: product.RatingCount,
			Image:       productImage.Path,
		})
	}
	return responseData
}

func DeleteProductById(id string) error {
	var product models.Products
	result := Db.Delete(&product, id)
	if result.RowsAffected < 1 {
		return errors.New("product not found")
	}
	return nil
}
