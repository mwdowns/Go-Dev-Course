package middleware

import (
	"mwdowns/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": "unauthorized"})
	}
	context.Set("userId", userId)
	context.Next()
}

func CheckUser(context *gin.Context, eventUserId int64) bool {
	userId, _ := context.Get("userId")
	if userId.(int64) != eventUserId {
		return false
	}
	return true
}
