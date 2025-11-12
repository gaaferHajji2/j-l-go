package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
