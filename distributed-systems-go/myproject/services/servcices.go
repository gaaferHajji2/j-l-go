package services

import (
	"context"
	"fmt"
	"myproject/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMealHandler(ctx context.Context, collection *mongo.Collection) *MealHandler {
	return &MealHandler{
		collection: collection,
		ctx:        ctx,
	}
}

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
func (handler *MealHandler) CreateMealHandler(c *gin.Context) {
	var meal models.Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	meal.ID = primitive.NewObjectID()
	meal.CreatedAt = time.Now()
	_, err := handler.collection.InsertOne(handler.ctx, &meal)
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
func (handler *MealHandler) GetAllMealsHandler(c *gin.Context) {
	cur, err := handler.collection.Find(handler.ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cur.Close(handler.ctx)

	meals := make([]models.Meal, 0)

	for cur.Next(handler.ctx) {
		var meal models.Meal
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
func (handler *MealHandler) UpdateMealHandler(c *gin.Context) {
	id := c.Param("id")

	var meal models.Meal

	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := handler.collection.UpdateOne(handler.ctx, bson.M{"_id": new_id}, bson.D{{Key: "$set", Value: bson.D{
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
func (handler *MealHandler) DeleteMealHandler(c *gin.Context) {
	id := c.Param("id")
	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := handler.collection.DeleteOne(handler.ctx, bson.M{"_id": new_id})

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
func (handler *MealHandler) SearchForMealByTag(c *gin.Context) {
	tag := c.Query("tag")

	cur, err := handler.collection.Find(handler.ctx, bson.M{"tags": tag})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	defer cur.Close(handler.ctx)

	meals := make([]models.Meal, 0)

	for cur.Next(handler.ctx) {
		var meal models.Meal
		cur.Decode(&meal)
		meals = append(meals, meal)
	}

	c.JSON(http.StatusOK, meals)
}
