package routes

import (
	"fmt"
	"mwdowns/rest-api/interfaces"
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

func updateObject(object interfaces.Updater, context *gin.Context, name string) {
	_, err := object.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update " + name, "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": name + " updated"})
}

func createObject(object interfaces.Saver, context *gin.Context, name string) {
	id, err := object.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save " + name, "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": name + " created", name: object})
	fmt.Printf("this is the object id: %v\n", id)
}

func deleteObject(object interfaces.Deleter, context *gin.Context, name string) {
	err := object.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete " + name, "error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"message": name + " deleted"})
}
