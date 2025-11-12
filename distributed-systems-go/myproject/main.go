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
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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

	mealsHandler = services.NewMealHandler(ctx, collection)
}

func main() {
	router := gin.Default()

	router.POST("/meals", mealsHandler.CreateMealHandler)
	router.GET("/meals", mealsHandler.GetAllMealsHandler)
	router.PUT("/meals/:id", mealsHandler.UpdateMealHandler)
	router.DELETE("/meals/:id", mealsHandler.DeleteMealHandler)
	router.GET("/meals/searchByTag", mealsHandler.SearchForMealByTag)

	router.Run()
}
