package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"os"
	"time"
)

const (
	Port = "3000"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML("<b> Hello API! </b>")
	})

	// API Routes
	apiRoutes := app.Party("/api", logThisMiddleware)
	{
		apiRoutes.Get("/", getAPIInfo)
		apiRoutes.Get("/todos", getTodos)
	}

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.Run(iris.Server(server))
}

func logThisMiddleware(ctx context.Context) {
	ctx.Application().Log("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func getAPIInfo(ctx context.Context) {
	ctx.JSON(map[string]interface{}{
		"version": "0.0.1",
	})
}

type TodoModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var todos = map[string]TodoModel{
	"1": TodoModel{Id: 1, Title: "Test1"},
	"2": TodoModel{Id: 2, Title: "Test2"},
	"3": TodoModel{Id: 3, Title: "Test3"},
}

func getTodos(ctx context.Context) {
	ctx.JSON(todos)
}
