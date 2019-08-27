package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type struct Person {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}