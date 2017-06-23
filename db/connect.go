package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo   *mgo.DialInfo

	//Collections
	Cinemas *mgo.Collection
	Halls   *mgo.Collection
	Movies  *mgo.Collection
	Seances *mgo.Collection
)

// Connect connects to mongodb
func Connect(MongoDBUrl string) {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	session.SetSafe(&mgo.Safe{})
	session.SetMode(mgo.Monotonic, true)

	fmt.Println("Connected to ", uri)

	Session = session
	Mongo = mongo

	Cinemas = session.DB(mongo.Database).C("cinemas")
	Halls = session.DB(mongo.Database).C("cinema_halls")
	Movies = session.DB(mongo.Database).C("movies")
	Seances = session.DB(mongo.Database).C("seances")

}
