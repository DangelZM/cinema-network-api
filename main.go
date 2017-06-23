package main

import (
	"os"
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/middlewares"
)

const (
	Port       = "3000"
	Version    = "0.0.1"
)

func init() {
	db.Connect("mongodb://localhost:27017/cinema_network")
}

func main() {

	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORS)

	registerRoutes(router)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Start listening on " + port)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
