package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/favicon.ico", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	router.GET("/events", showEvents)
	router.GET("/events/:id", showEvent)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
}
