package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string        `json:"title" binding:"required" bson:"title"`
	Year     int           `json:"year" binding:"required" bson:"year"`
	Duration int           `json:"duration" binding:"required" bson:"duration"`
}
