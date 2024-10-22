package controllers

import (
	"context"
	"persogo/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateEmployee(collection *mongo.Collection, employee models.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
}
