package cinemas

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/models"
	"time"
)

func Create(c *gin.Context) {
	cinema := models.Cinema{}

	if err := c.Bind(&cinema); err != nil {
		c.Error(err)
		return
	}

	if err := db.Cinemas.Insert(cinema); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, cinema)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := db.Cinemas.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func GetList(c *gin.Context) {
	cinemas := []models.Cinema{}

	if err := db.Cinemas.Find(nil).All(&cinemas); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, cinemas)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	cinema := models.CinemaWithHalls{}

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": bson.ObjectIdHex(id)}},
		bson.M{"$lookup": bson.M{
			"from":         "halls",
			"localField":   "_id",
			"foreignField": "cinema_id",
			"as":           "halls",
		}},
	}

	if err := db.Cinemas.Pipe(pipeline).One(&cinema); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cinema)
}

func GetHalls(c *gin.Context) {
	id := c.Params.ByName("id")

	halls := []models.Hall{}

	if err := db.Halls.Find(bson.M{"cinema_id": bson.ObjectIdHex(id)}).All(&halls); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, halls)
}

func GetSeancesByDate(c *gin.Context) {
	id := c.Params.ByName("id")
	date := c.Params.ByName("date")

	if date == "today" {
		date = time.Now().Format("2006-01-02")
	}

	dayStart, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.Error(errors.New("Not valid date"))
		return
	}

	dayEnd := dayStart.Add(time.Hour * 24)

	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"$and": []bson.M{
				bson.M{"cinema_id": bson.ObjectIdHex(id)},
				bson.M{"$and": []bson.M{
					{"start_time": bson.M{
						"$gte": dayStart,
					}},
					{"end_time": bson.M{
						"$lte": dayEnd,
					}},
				}},
			},
		}},
		bson.M{"$lookup": bson.M{
			"from":         "movies",
			"localField":   "movie_id",
			"foreignField": "_id",
			"as":           "movies",
		}},
		bson.M{"$addFields": bson.M{
			"movie": bson.M{
				"$arrayElemAt": []interface{}{
					"$movies",
					0,
				},
			},
			"movie_id": "",
		}},
		bson.M{"$sort": bson.M{
			"start_time": 1,
		}},
	}

	seances := []models.SeanceWithMovie{}

	if err := db.Seances.Pipe(pipeline).All(&seances); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, seances)
}
