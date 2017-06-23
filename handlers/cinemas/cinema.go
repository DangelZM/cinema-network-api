package cinemas

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/models"
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

func GetList(c *gin.Context) {
	cinemas := []models.Cinema{}

	if err := db.Cinemas.Find(nil).All(&cinemas); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, cinemas)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	cinema := models.Cinema{}

	if !bson.IsObjectIdHex(id) {
		c.Error(errors.New("Not valid id"))
		return
	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Error(errors.New("Not valid id"))
		return
	} else if err := db.Cinemas.Find(bson.M{"_id": obj}).One(&cinema); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, cinema)
}

func GetHalls(c *gin.Context) {
	id := c.Params.ByName("id")

	halls := []models.Hall{}

	if !bson.IsObjectIdHex(id) {
		c.Error(errors.New("Not valid id"))
		return
	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Error(errors.New("Not valid id"))
		return
	} else if err := db.Halls.Find(bson.M{"cinema_id": obj}).All(&halls); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, halls)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.Error(errors.New("Not valid id"))
		return
	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Error(errors.New("Not valid id"))
		return
	} else if err := db.Cinemas.Remove(bson.M{"_id": obj}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
