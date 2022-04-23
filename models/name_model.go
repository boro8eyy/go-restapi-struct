package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NameModel struct {
	Id primitive.ObjectID `bson:"id,omitempty"`
}