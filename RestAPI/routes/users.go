package routes

import (
	"fmt"
	"mwdowns/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	// takes in from post and turns it into User
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	id, err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
	fmt.Printf("this is the user id: %v\n", id)
}
