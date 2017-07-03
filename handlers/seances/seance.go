package seances

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/models"
	"time"
)

const PauseDuration = 15 * time.Minute

func checkSeanceTime(seance models.Seance) (bool, error) {
	checkStart := seance.Start.Add(-time.Duration(PauseDuration))
	checkEnd := seance.End.Add(time.Duration(PauseDuration))

	condition := bson.M{
		"hall_id": seance.HallId,
		"$or": []bson.M{
			{
				"$and": []bson.M{
					{"start_time": bson.M{
						"$lte": checkStart,
					}},
					{"end_time": bson.M{
						"$gte": checkStart,
					}},
				},
			},
			{
				"$and": []bson.M{
					{"start_time": bson.M{
						"$lte": checkEnd,
					}},
					{"end_time": bson.M{
						"$gte": checkEnd,
					}},
				},
			},
			{
				"$and": []bson.M{
					{"start_time": bson.M{
						"$lte": checkStart,
					}},
					{"end_time": bson.M{
						"$gte": checkEnd,
					}},
				},
			},
			{
				"$and": []bson.M{
					{"start_time": bson.M{
						"$gte": checkStart,
					}},
					{"end_time": bson.M{
						"$lte": checkEnd,
					}},
				},
			},
		},
	}

	existSeances := []models.Seance{}
	if err := db.Seances.Find(condition).All(&existSeances); err != nil {
		return false, err
	}

	if len(existSeances) > 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func Create(c *gin.Context) {
	seance := models.Seance{}

	if err := c.Bind(&seance); err != nil {
		c.Error(err)
		return
	}

	movie := models.Movie{}
	if err := db.Movies.FindId(seance.MovieId).One(&movie); err != nil {
		c.Error(err)
		return
	}

	seance.End = seance.Start.Add(time.Duration(movie.Duration) * time.Minute)

	check, err := checkSeanceTime(seance)

	if err != nil {
		c.Error(err)
		return
	}

	if check {
		if err := db.Seances.Insert(seance); err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, seance)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Seance exist for selected time",
		})
	}

}

func GetList(c *gin.Context) {
	seances := []models.Seance{}

	if err := db.Seances.Find(nil).All(&seances); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, seances)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	seance := models.Seance{}

	if err := db.Seances.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&seance); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, seance)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := db.Seances.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
