package main

import (
	"net/http"
	"strconv"

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

	r2 := r.Group("/profile")
	{
		r2.GET("/image/:id", getImageById)
		r2.GET("/:username", getProfileByUsername)
	}

	r.Run(":8001")
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
