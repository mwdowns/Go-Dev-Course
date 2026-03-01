package routes

import (
	"mwdowns/rest-api/models"
	"mwdowns/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const usersName = "users"

func createUser(context *gin.Context) {
	// takes in from post and turns it into User
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	createObject(user, context, usersName)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid user", "error": "wrong email or password"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user logged in", "token": token})
}
