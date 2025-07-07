package middlewares

import (
	"net/http"

	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you must be logged in to create an event."})
		return
	}

	user_id, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token, please log in again."})
		return
	}

	context.Set("user_id", user_id)

	context.Next()
}
