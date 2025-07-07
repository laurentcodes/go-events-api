package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	user_id := context.GetInt64("user_id")
	event_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id."})
		return
	}

	event, err := models.GetEvent(event_id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, try again later."})
		return
	}

	err = event.Register(user_id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user registered for event successfully."})
}

func cancelRegistration(context *gin.Context) {
	user_id := context.GetInt64("user_id")
	event_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = event_id

	err = event.CancelRegistration(user_id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration cancelled successfully."})
}
