package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Cinema struct {
	Id    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" binding:"required" bson:"title"`
}

type CinemaWithHalls struct {
	Cinema `bson:",inline"`
	Halls  []Hall `json:"halls" bson:"halls"`
}
