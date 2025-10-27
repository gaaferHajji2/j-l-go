package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default();

	router.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			// the comma at the end is required
			"msg": "Hello Jafar Loka World",
		})
	})

	router.Run(":5000")
}