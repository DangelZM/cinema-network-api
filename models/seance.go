package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Seance struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CinemaId bson.ObjectId `json:"cinema_id" binding:"required" bson:"cinema_id"`
	HallId   bson.ObjectId `json:"hall_id" binding:"required" bson:"hall_id"`
	MovieId  bson.ObjectId `json:"movie_id,omitempty" binding:"required" bson:"movie_id"`
	Start    time.Time     `json:"start_time" binding:"required" bson:"start_time"`
	End      time.Time     `json:"end_time,omitempty" bson:"end_time"`
}

type SeanceWithMovie struct {
	Seance `bson:",inline"`
	Movie  Movie `json:"movie" bson:"movie"`
}
