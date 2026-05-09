package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// When we don't send stock data, the default value is 0
type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock"`
	Price float32 `json:"price" binding:"required"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error in reading env file")
	}

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello From Jafar Loka",
		})
	})

	r1 := r.Group("/api")
	{
		r1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello From Jafar Loka - GET",
			})
		})

		r1.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello From Jafar Loka - POST",
			})
		})

		r1.PUT("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello From Jafar Loka - PUT",
			})
		})

		r1.DELETE("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello From Jafar Loka - DELETE",
			})
		})
	}

	r2 := r.Group("/profile")
	{
		r2.GET("/image/:id", getImageById)
		r2.GET("/:username", getProfileByUsername)
		r2.GET("/", showData)
	}

	r.POST("/handleProduct", handleProduct)
	r.POST("/handleProducts", handleProducts)

	fmt.Println("The port is: ", os.Getenv("PORT"))
	fmt.Println("The host is: ", os.Getenv("HOST"))

	r.Run(":" + os.Getenv("PORT"))
}

func handleProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid JSON: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func handleProducts(c *gin.Context) {
	var products []Product
	if err := c.BindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid JSON: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func getImageById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	// id, err = strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Invalid value for id: " + err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":  id,
		"msg": "Image OK",
	})
}

func getProfileByUsername(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"username": c.Param("username"),
		"msg":      "Profile Ok",
	})
}

func showData(c *gin.Context) {
	// http://localhost:8001/profile/?id=2&jobs[0]=IT&jobs=developer&jobs=designer&jobs=manager
	// &data[id]=100&data[name]=JLoka&data[type]=Employee&age=26
	c.JSON(http.StatusOK, gin.H{
		"age":  c.DefaultQuery("age", "10"),
		"id":   c.Query("id"),
		"jobs": c.QueryArray("jobs"),
		"data": c.QueryMap("data"),
	})
}
