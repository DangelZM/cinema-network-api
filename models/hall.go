package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Hall struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CinemaId bson.ObjectId `json:"cinema_id" binding:"required" bson:"cinema_id"`
	Title    string        `json:"title" binding:"required" bson:"title"`
	Seating  int           `json:"seating" binding:"required" bson:"seating"`
}
