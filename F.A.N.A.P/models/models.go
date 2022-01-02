package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rectangle struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	X      float64            `json:"x,omitempty"`
	Y      float64            `json:"y,omitempty"`
	Width  float64            `json:"width,omitempty"`
	Height float64            `json:"height,omitempty"`
	Time   string             `json:"time,omitempty"`
}
