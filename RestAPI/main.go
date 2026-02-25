package main

import (
	"fmt"
	"mwdowns/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//db.InitDB()
	router := gin.Default()
	router.GET("/favicon.ico", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	router.GET("/events", showEvents)
	router.POST("/events", createEvent)

	router.Run(":8080")
}

func showEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
	fmt.Printf("this is the first event id: %v\n", events[0].ID)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	id, err := event.Save("events")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
	fmt.Printf("this is the event id: %v\n", id)
}
