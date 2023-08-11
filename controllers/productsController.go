package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ilyasbabu/go-ecommerce/models"
	"github.com/ilyasbabu/go-ecommerce/services"
)

func Ping(c *gin.Context) {
	c.String(200, "Pong")
}

func CreateProductAdmin(c *gin.Context) {
	r := Response()
	validatedData, err := services.ValidateProductCreate(c)
	if err != nil {
		r.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	err = services.CreateProduct(validatedData, c)
	if err != nil {
		r.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	r.Status = "SUCCESS"
	r.Message = "Product Created Successfully"
	c.JSON(http.StatusOK, r)
}

func GetProduct(c *gin.Context) {
	r := Response()

	data, err := services.GetProductByID(c.Param("id"))
	if err != nil {
		r.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	r.Status = "SUCCESS"
	r.Data = data
	r.Message = "Product Fetched Successfully"
	c.JSON(http.StatusOK, r)
}

func GetProducts(c *gin.Context) {
	r := Response()
	data := services.GetProducts()
	r.Status = "SUCCESS"
	r.Data = data
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
	if err := services.DeleteProductById(c.Param("id")); err != nil {
		r.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	r.Status = "SUCCESS"
	r.Message = "Product Updated Successfully"
	c.JSON(http.StatusOK, r)
}
