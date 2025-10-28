package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName xml.Name `xml:"PersonObject"`
	// here the data and attr must start with capital letter
	FullName string `xml:"FullName"`
	Age      int16  `xml:"Age,attr"`
}

func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello Jafar Loka World",
	})
}

func HelloWorldNameHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"msg": "Hello To JLoka World: " + name,
	})
}

func FullNameHandler(c *gin.Context) {
	c.XML(200, Person{FullName: "Jafar Loka-01", Age: 25})
}

func main() {
	router := gin.Default()

	router.GET("/", HelloWorldHandler)

	router.GET("/:name", HelloWorldNameHandler)

	router.GET("/fullName", FullNameHandler)

	router.Run(":5000")
}
