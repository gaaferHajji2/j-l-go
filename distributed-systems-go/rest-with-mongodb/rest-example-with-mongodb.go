package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Meal struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	CreatedAt    time.Time `json:"createdAt"`
}

var meals []Meal

func CreateMealHandler(c *gin.Context) {
	var meal Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	meal.ID = xid.New().String()
	meal.CreatedAt = time.Now()
	meals = append(meals, meal)
	c.JSON(http.StatusCreated, meal)
}

func GetAllMealsHandler(c *gin.Context) {

	// if len(meals) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "No Data Found",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, meals)
}

func UpdateMealHandler(c *gin.Context) {
	id := c.Param("id")

	var meal Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	for i := 0; i < len(meals); i++ {
		if meals[i].ID == id {
			meal.CreatedAt = meals[i].CreatedAt
			meal.ID = id

			meals[i] = meal
			c.JSON(http.StatusOK, meal)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"err": "No Meal Found",
	})
}

func main() {
	router := gin.Default()

	meals = make([]Meal, 0)

	router.POST("/meals", CreateMealHandler)

	router.GET("/meals", GetAllMealsHandler)

	router.PUT("/meals/:id", UpdateMealHandler)

	router.Run()
}
