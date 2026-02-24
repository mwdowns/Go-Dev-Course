package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
	uuid        uuid.UUID
}

var events = []Event{}

func (e Event) Save() {
	// store in DB
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}

func CreateEvent(context *gin.Context) {}
