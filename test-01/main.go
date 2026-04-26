package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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

	r.Run(":8001")
}
