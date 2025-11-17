// Package main My Simple Meal API
//
// The purpose of this application is to demonstrate go-swagger
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/redis/go-redis/v9"

	services "myproject/services"
)

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection
var mealsHandler *services.MealHandler

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Failed to connect to mongodb", err)
	}
	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	log.Println("Successfully connected to mongodb")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	mealsHandler = services.NewMealHandler(ctx, collection, redisClient)

	status := redisClient.Ping(ctx)
	fmt.Println("The status is: ", status)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != os.Getenv("X_API_KEY") {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()

	authMealsGroup := router.Group("/")

	authMealsGroup.Use(AuthMiddleware())
	{
		authMealsGroup.POST("/meals", mealsHandler.CreateMealHandler)
		authMealsGroup.PUT("/meals/:id", mealsHandler.UpdateMealHandler)
		authMealsGroup.DELETE("/meals/:id", mealsHandler.DeleteMealHandler)
	}

	router.GET("/meals", mealsHandler.GetAllMealsHandler)

	router.GET("/meals/searchByTag", mealsHandler.SearchForMealByTag)

	router.Run()
}
