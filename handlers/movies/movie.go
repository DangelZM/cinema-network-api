package movies

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/models"
)

func Create(c *gin.Context) {
	movie := models.Movie{}

	if err := c.Bind(&movie); err != nil {
		c.Error(err)
		return
	}

	if err := db.Movies.Insert(movie); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func GetList(c *gin.Context) {
	movies := []models.Movie{}

	if err := db.Movies.Find(nil).All(&movies); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, movies)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	movie := models.Movie{}

	if err := db.Movies.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&movie); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, movie)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := db.Movies.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
