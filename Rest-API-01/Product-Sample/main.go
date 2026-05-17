package main

import (
	"jloka/jloka/controllers"
	"jloka/jloka/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var products []models.Product
	var productRepo controllers.ProductRepo = *controllers.Init(&products)

	r.POST("/products", productRepo.CreateProduct)
	r.GET("/products", productRepo.GetProducts)
	r.GET("/products/:id", productRepo.GetProductById)
	r.PUT("/products/:id", productRepo.UpdateProductById)
	r.DELETE("/products/:id", productRepo.DeleteProduct)

	r.Run()
}
