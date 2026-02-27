package routes

import (
	"fmt"
	"mwdowns/rest-api/middleware"
	"mwdowns/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
	fmt.Printf("this is the first event id: %v\n", events[0].Uuid)
}

func showEvent(context *gin.Context) {
	id := context.Param("id")
	e, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "got event", "event": e})
	fmt.Printf("this is the event id: %v\n", e.Location)
}

func createEvent(context *gin.Context) {
	// takes in from post and turns it into Event
	var event models.Event
	err := context.ShouldBindJSON(&event)
	// the token has the userId for who created the event, so we attach it to the event
	userId, ok := context.Get("userId")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get userId", "error": err.Error()})
	}
	event.UserID = userId.(string)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	id, err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
	fmt.Printf("this is the event id: %v\n", id)
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	_, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event", "error": err.Error()})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	eventUserId := updatedEvent.UserID
	if !middleware.CheckUser(context, eventUserId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
		return
	}

	updatedEvent.Uuid = id
	id, err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated", "event_id": id})
}

func removeEvent(context *gin.Context) {
	id := context.Param("id")
	e, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could event to delete", "error": err.Error()})
		return
	}
	if !middleware.CheckUser(context, e.UserID) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
		return
	}
	err = e.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted", "event_id": id})
}
