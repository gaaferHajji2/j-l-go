package controllers

import (
	"jloka/jloka/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

// Create Product For Controller
func (productRepo *ProductRepo) CreateProduct(c *gin.Context) {
	var product *models.Product

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Product Request: " + err.Error(),
		})
		return
	}
	if err := product.CreateProduct(productRepo.Products); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "ISE: " + err.Error(),
		})
	}

	c.JSON(http.StatusCreated, product)
}
