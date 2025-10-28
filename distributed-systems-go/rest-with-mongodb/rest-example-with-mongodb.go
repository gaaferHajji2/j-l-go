package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Meal struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	CreatedAt    time.Time `json:"createdAt"`
}

func main() {
	router := gin.Default()

	router.Run()
}
