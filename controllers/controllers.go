package controllers

import (
	"context"
	"log"
	"persogo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateEmployee(collection *mongo.Collection, employee models.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	employee.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, employee)
	if err != nil {
		log.Println("Ошибка при создании сотрудника!", err)
		return err
	}
	log.Println("Сотрудник добавлен!")
	return nil
}

func GetEmployee(collection *mongo.Collection) ([]models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var employees []models.Employee
	if err = cursor.All(ctx, &employees); err != nil {
		return nil, err
	}
	return employees, nil
}
