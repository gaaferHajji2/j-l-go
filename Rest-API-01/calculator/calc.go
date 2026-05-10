package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Numeric struct {
	Num1 float64 `json:"num1" binding:"required"`
	Num2 float64 `json:"num2" binding:"required"`
}

func Add(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result is: ": body.Num1 + body.Num2,
	})
}

func Subtract(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result is: ": body.Num1 - body.Num2,
	})
}

func Multiply(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result is: ": body.Num1 * body.Num2,
	})
}

func Divide(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result is: ": body.Num1 / body.Num2,
	})
}
