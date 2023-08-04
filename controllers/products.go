package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ilyasbabu/go-ecommerce/models"
)

func Ping(c *gin.Context) {
	c.String(200, "Pong")
}

func CreateProductAdmin(c *gin.Context) {
	r := Response()
	name := c.PostForm("name")
	if name == "" {
		r.Message = "Name must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	description := c.PostForm("description")
	if description == "" {
		r.Message = "Description must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	price := c.PostForm("price")
	if price == "" {
		r.Message = "Price must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		r.Message = "Price Invalid"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	product := models.Products{Name: name, Description: description, Price: floatPrice}
	result := Db.Create(&product)
	if result.Error != nil {
		r.Message = result.Error.Error()
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			r.Message = "Cannot Update! Product with same name(" + product.Name + ") already exists"
		}
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	r.Status = "SUCCESS"
	r.Data = gin.H{
		"id":   product.ID,
		"slug": product.Slug,
	}
	r.Message = "Product Created Successfully"
	c.JSON(http.StatusOK, r)
}

func GetProduct(c *gin.Context) {
	r := Response()
	var product models.Products

	id := c.Param("id")
	err := Db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.Message = "Product not found"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	data := struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createAt"`
		Slug        string    `json:"slug"`
		Stock       int       `json:"stock"`
		AvgRating   float64   `json:"avgRating"`
		RatingCount int       `json:"ratingCount"`
	}{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		Slug:        product.Slug,
		Stock:       product.Stock,
		AvgRating:   product.AvgRating,
		RatingCount: product.RatingCount,
	}
	r.Status = "SUCCESS"
	r.Data = data
	r.Message = "Product Fetched Successfully"
	c.JSON(http.StatusOK, r)
}

func GetProducts(c *gin.Context) {
	r := Response()
	var products []models.Products

	result := Db.Find(&products)
	if result.RowsAffected == 0 {
		r.Message = "Product not found"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	type productResponsese struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createAt"`
		Slug        string    `json:"slug"`
		Stock       int       `json:"stock"`
		AvgRating   float64   `json:"avgRating"`
		RatingCount int       `json:"ratingCount"`
	}
	var responseData []productResponsese

	for _, product := range products {
		responseData = append(responseData, productResponsese{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			Slug:        product.Slug,
			Stock:       product.Stock,
			AvgRating:   product.AvgRating,
			RatingCount: product.RatingCount,
		})
	}
	r.Status = "SUCCESS"
	r.Data = responseData
	r.Message = "Products Fetched Successfully"
	c.JSON(http.StatusOK, r)
}

func UpdateProduct(c *gin.Context) {
	r := Response()
	var product models.Products

	id := c.Param("id")
	err := Db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.Message = "Product not found"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	name := c.PostForm("name")
	if name == "" {
		r.Message = "Name must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	description := c.PostForm("description")
	if description == "" {
		r.Message = "Description must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	price := c.PostForm("price")
	if price == "" {
		r.Message = "Price must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		r.Message = "Price Invalid"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	product.Name = name
	product.Description = description
	product.Price = floatPrice

	result := Db.Save(&product)
	if result.Error != nil {
		r.Message = result.Error.Error()
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			r.Message = "Cannot Update! Product with same name(" + product.Name + ") already exists"
		}
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	r.Status = "SUCCESS"
	r.Data = "data"
	r.Message = "Product Updated Successfully"
	c.JSON(http.StatusOK, r)
}

func DeleteProduct(c *gin.Context) {
	r := Response()
	var product models.Products

	id := c.Param("id")
	result := Db.Delete(&product, id)
	if result.RowsAffected < 1 {
		r.Message = "Product not found"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	r.Status = "SUCCESS"
	r.Data = "data"
	r.Message = "Product Updated Successfully"
	c.JSON(http.StatusOK, r)
}
