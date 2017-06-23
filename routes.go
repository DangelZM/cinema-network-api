package main

import (
	"github.com/gin-gonic/gin"

	cinemasHandler "github.com/dangelzm/cinema-network-api/handlers/cinemas"
	hallsHandler "github.com/dangelzm/cinema-network-api/handlers/halls"
	moviesHandler "github.com/dangelzm/cinema-network-api/handlers/movies"
	seancesHandler "github.com/dangelzm/cinema-network-api/handlers/seances"
)

func registerRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) { c.String(200, "OK") })

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": Version,
			})
		})

		cinemas := api.Group("/cinemas")
		{
			cinemas.POST("/", cinemasHandler.Create)
			cinemas.GET("/", cinemasHandler.GetList)
			cinemas.GET("/:id", cinemasHandler.GetOne)
			cinemas.DELETE("/:id", cinemasHandler.Delete)
			cinemas.GET("/:id/halls", cinemasHandler.GetHalls)
		}

		halls := api.Group("/halls")
		{
			halls.POST("/", hallsHandler.Create)
			halls.GET("/", hallsHandler.GetList)
			halls.GET("/:id", hallsHandler.GetOne)
			halls.DELETE("/:id", hallsHandler.Delete)
		}

		movie := api.Group("/movies")
		{
			movie.POST("/", moviesHandler.Create)
			movie.GET("/", moviesHandler.GetList)
			movie.GET("/:id", moviesHandler.GetOne)
			movie.DELETE("/:id", moviesHandler.Delete)
		}

		seance := api.Group("/seances")
		{
			seance.POST("/", seancesHandler.Create)
			seance.GET("/", seancesHandler.GetList)
			seance.GET("/:id", seancesHandler.GetOne)
			seance.DELETE("/:id", seancesHandler.Delete)
		}
	}

	router.NoRoute()

	return
}
