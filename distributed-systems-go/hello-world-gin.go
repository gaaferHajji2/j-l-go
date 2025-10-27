package main

import "github.com/gin-gonic/gin"

func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello Jafar Loka World",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", HelloWorldHandler)

	router.Run(":5000")
}
