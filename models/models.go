package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"first_name,omitempty"`
	LastName    string             `bson:"last_name,omitempty"`
	Position    string             `bson:"position,omitempty"`
	Departament string             `bson:"departament,omiempty"`
}
