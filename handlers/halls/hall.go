package halls

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/models"
)

func Create(c *gin.Context) {
	Hall := models.Hall{}

	if err := c.Bind(&Hall); err != nil {
		c.Error(err)
		return
	}

	if err := db.Halls.Insert(Hall); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, Hall)
}

func GetList(c *gin.Context) {
	halls := []models.Hall{}

	if err := db.Halls.Find(nil).All(&halls); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, halls)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	Hall := models.Hall{}

	if err := db.Halls.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&Hall); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, Hall)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := db.Halls.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
