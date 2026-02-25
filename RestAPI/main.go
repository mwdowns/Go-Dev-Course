package main

import (
	"mwdowns/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/favicon.ico", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	router.GET("/events", showEvents)
	router.POST("/events", createEvent)

	router.Run(":8080")
}

func showEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
