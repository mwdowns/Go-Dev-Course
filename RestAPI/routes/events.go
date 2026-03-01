package routes

import (
	"fmt"
	"mwdowns/rest-api/middleware"
	"mwdowns/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const eventsName = "event"

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
	parsedId, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse event id", "error": err.Error()})
	}
	e, err := models.GetEvent(parsedId)
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
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}
	// the token has the userId for who created the event, so we attach it to the event
	userId, ok := context.Get("userId")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get userId from context", "error": err.Error()})
	}
	event.UserID = userId.(int64)
	createObject(event, context, eventsName)
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id", "error": err.Error()})
	}
	event, err := models.GetEvent(parsedUuid)
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

	updatedEvent.Uuid = parsedUuid
	eventUserId := event.UserID
	if !middleware.CheckUser(context, eventUserId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
		return
	}
	updateObject(updatedEvent, context, eventsName)
}

func removeEvent(context *gin.Context) {
	id := context.Param("id")
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse id", "error": err.Error()})
	}
	e, err := models.GetEvent(parsedUuid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could event to delete", "error": err.Error()})
		return
	}
	if !middleware.CheckUser(context, e.UserID) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
		return
	}
	deleteObject(e, context, eventsName)
}
