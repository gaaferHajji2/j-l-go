package services

import (
	"context"

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
