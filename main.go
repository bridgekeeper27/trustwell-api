package main

import (
	"github.com/bridgekeeper27/trustwell-api/database"
	"github.com/bridgekeeper27/trustwell-api/handlers"
	"github.com/bridgekeeper27/trustwell-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.Use(middleware.AuthRequired())

	r.POST("/events", handlers.CreateEvent)
	r.GET("/events", handlers.ListEvents)
	r.DELETE("/events/:id", handlers.DeleteEvent)
	r.GET("/events/:id", handlers.GetEvent)

	r.Run()
}
