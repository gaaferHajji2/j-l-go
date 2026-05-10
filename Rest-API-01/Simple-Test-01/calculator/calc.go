package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Numeric struct {
	Num1 float64 `json:"num1" binding:"required"`
	Num2 float64 `json:"num2" binding:"required"`
	Res  float64 `json:"res"`
}

func Add(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}
	body.Res = body.Num1 + body.Num2

	c.JSON(http.StatusOK, body)
}

func Subtract(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}
	body.Res = body.Num1 - body.Num2

	c.JSON(http.StatusOK, body)
}

func Multiply(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}
	body.Res = body.Num1 * body.Num2

	c.JSON(http.StatusOK, body)
}

func Divide(c *gin.Context) {
	var body Numeric
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	body.Res = body.Num1 / body.Num2

	c.JSON(http.StatusOK, body)
}
