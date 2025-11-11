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
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client

// Meal represents a object in the system.
//
//	swagger:model meal
type Meal struct {
	// ID of the Meal
	//	required: true
	ID primitive.ObjectID `json:"id" bson:"_id"`
	// Name of the Meal
	//	required: true
	Name string `json:"name" bson:"name"`
	// Tags of the meal
	//	required: true
	Tags []string `json:"tags" bson:"tags"`
	// Ingredients of meal
	//	required: true
	Ingredients []string `json:"ingredients" bson:"ingredients"`
	// Instructions to create meal
	//	required: true
	Instructions []string `json:"instructions" bson:"instructions"`
	// When we create the record
	//	required: true
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

var meals []Meal

// swagger:operation POST /meals meals CreateMealHandler
// Create meal
// ---
//
//		parameters:
//		- name: meal
//		  in: body
//		  description: The body of request
//		  required: true
//		  schema:
//	      "$ref": "#/definitions/meal"
//		consumes:
//		- application/json
//		produces:
//		- application/json
//		responses:
//			'201':
//				description: Meal Created Successfully
//				schema:
//	      		"$ref": "#/definitions/meal"
//			'400':
//				description: Bad Request
func CreateMealHandler(c *gin.Context) {
	var meal Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	meal.ID = primitive.NewObjectID()
	meal.CreatedAt = time.Now()
	_, err := collection.InsertOne(ctx, &meal)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, meal)
}

//	swagger:operation GET /meals meals listOfMeals
//
// Returns list of meals
// ---
//
//	produces:
//	- application/json
//	responses:
//		'200':
//			description: list of meals
func GetAllMealsHandler(c *gin.Context) {
	// if len(meals) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "No Data Found",
	// 	})
	// 	return
	// }

	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("recipes")

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cur.Close(ctx)

	meals := make([]Meal, 0)

	for cur.Next(ctx) {
		var meal Meal
		cur.Decode(&meal)
		meals = append(meals, meal)
	}

	c.JSON(http.StatusOK, meals)

}

// swagger:operation PUT /meals/{id} meals UpdateMealById
// Update meal by id if exists only
// ---
//
//	parameters:
//	- name: id
//	  in: path
//	  description: id of the meal
//	  required: true
//	  type: string
//	produces:
//	- application/json
//	responses:
//		'200':
//			description: Meal Updated Successfully
//		'400':
//			description: Bad Request
//		'404':
//			description: Meal not found
func UpdateMealHandler(c *gin.Context) {
	id := c.Param("id")

	var meal Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := collection.UpdateOne(ctx, bson.M{"_id": new_id}, bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: meal.Name},
		{Key: "instructions", Value: meal.Instructions},
		{Key: "ingredients", Value: meal.Ingredients},
		{Key: "tags", Value: meal.Tags},
	}}})

	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

//	swagger:operation DELETE /meals/{id} meals DeleteMealHandler
//
// Delete meal by id if exists
// ---
//
//	parameters:
//	- name: id
//	  in: path
//	  description: id of the meal
//	  required: true
//	  type: string
//	responses:
//		'204':
//			description: The meal deleted successfully
//		'404':
//			description: meal deleted successfully
func DeleteMealHandler(c *gin.Context) {
	id := c.Param("id")
	new_id, _ := primitive.ObjectIDFromHex(id)

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	res, err := collection.DeleteOne(ctx, bson.M{"_id": new_id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	fmt.Println("res: ", res)
	c.JSON(http.StatusNoContent, gin.H{})
}

//	swagger:operation GET /meals/searchByTag meals SearchForMealByTag
//
// Returns list of meals that contains the tag
// ---
//	parameters:
//	- name: tag
//	  in: query
//	  description: tag of the meal
//	  required: true
//	  type: string

// produces:
// - application/json
// responses:
//
//	'200':
//		description: list of meals that contains the tag
func SearchForMealByTag(c *gin.Context) {
	tag := c.Query("tag")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	cur, err := collection.Find(ctx, bson.M{"tags": tag})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	meals := make([]Meal, 0)

	for cur.Next(ctx) {
		var meal Meal
		cur.Decode(&meal)
		meals = append(meals, meal)
	}

	c.JSON(http.StatusOK, meals)
}

func initFunc() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Failed to connect to mongodb", err)
	}

	log.Println("Successfully connected to mongodb")
}

func main() {
	router := gin.Default()
	meals = make([]Meal, 0)
	initFunc()

	router.POST("/meals", CreateMealHandler)
	router.GET("/meals", GetAllMealsHandler)
	router.PUT("/meals/:id", UpdateMealHandler)
	router.DELETE("/meals/:id", DeleteMealHandler)
	router.GET("/meals/searchByTag", SearchForMealByTag)

	router.Run()
}
