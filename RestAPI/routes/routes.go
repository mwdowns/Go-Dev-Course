package routes

import (
	"mwdowns/rest-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/favicon.ico", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	router.GET("/events", showEvents)
	router.GET("/events/:id", showEvent)
	router.POST("/signup", createUser)
	router.POST("/login", login)

	// authenticated routes
	authenticated := router.Group("/")
	// this sets the authenticate as the middleware for this group
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", removeEvent)
}
