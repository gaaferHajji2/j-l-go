package controllers

import (
	"fmt"
	"jloka/jloka/models"
	"net/http"
	"strconv"

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

func (productRepo *ProductRepo) DeleteProduct(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Param("id"))
	for _, product := range *productRepo.Products {
		if product.Id == Id {
			product.DeleteProductById(productRepo.Products)
			c.JSON(http.StatusNoContent, gin.H{
				"msg": "Data Deleted Successfully",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "No Data Found",
	})
}

func (productRepo *ProductRepo) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, productRepo.Products)
}

func (productRepo *ProductRepo) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	product := models.GetProductById(productRepo.Products, id)
	if product.Name != "" {
		fmt.Println("The Product Name is: ", product.Name)
		c.JSON(http.StatusNotFound, gin.H{})
	}
	c.JSON(http.StatusOK, product)
}

func (productRepo *ProductRepo) UpdateProductById(c *gin.Context) {
	id, err := strconv.Atoi("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	tempProduct := models.GetProductById(productRepo.Products, id)

	if tempProduct.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Product Not Found For Updating",
		})
		return
	}

	product := models.UpdateProductById(productRepo.Products, tempProduct)

	if product.Name != "" {
		fmt.Println("The Product Name is: ", product.Name)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, product)
}
