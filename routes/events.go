package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id."})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data."})
		return
	}

	user_id := context.GetInt64("user_id")

	event.UserID = user_id

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event, try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id."})
		return
	}

	user_id := context.GetInt64("user_id")
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, try again later."})
		return
	}

	if event.UserID != user_id {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data."})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated!"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id."})
		return
	}

	user_id := context.GetInt64("user_id")
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event, try again later."})
		return
	}

	if event.UserID != user_id {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to delete this event."})
		return
	}

	err = models.Delete(event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event, try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted!"})
}
