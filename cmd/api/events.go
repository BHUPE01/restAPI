package main

import (
	"net/http"
	"restapi/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJson(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faildes to crete eeen"})
		return
	}

	c.JSON(http.StatusCreated, event)

}

func (app *application) getAllEvents(c *gin.Context) {
	event, err := app.models.Events.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "filed to retreive events"})

	}

	c.JSON(http.StatusOK, event)
}

func (app *application) getEvent(c *gin.Context) {

	event, err := app.models.Events.Get(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": "invlaid event id"})
		return

	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failedto retrive event"})
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": "invlaid event id"})
		return
	}

	existingEvent, err := app.models.Events.Get(id)

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failedto retrive event"})
	}

	if err := c.ShouldBindJSON(app.updateEvent()); err != nil {
		c.JSON(http.StatusBindRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.Id = idapp
	if err != app.models.Events.Update(updateEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erroe": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erorr": "invalid event id"})

	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete event"})

	}
	c.JSON(http.StatusNoContent, nil)
}


func (app *application)addAttendeeToEvent(c *gin.Context){
	eventId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invlid event id"})
		return 
	}

	userId,err:=strconv.Atoi(c.Param("userId"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invlid event id"})
		return 
	}

	event,err:=app.models.Events.Get(eventId)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"invlid event id"})
		return 
	}

	if event==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"invlid event id"})
		
	}

	userToAdd,err:=app.models.Events.Get(userId)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error";"Failde to rtruvr"})
		return 
	}

	if userToAdd==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"User not found"})
	}

	existingAttendee,err:=app.models.Events.GetByEventAndAttendee(userId)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error";"Failde to rtruvrive attendee"})
		return 
	}

	if existingAttendee!=nil{
		c.JSON(http.StatusConflict,gin.H{"error":"attend e alredy exits"})
	}

	attendee:=database.Attendee{
		EventId:event.Id,
		UserId:userToAdd.Id,
	}

	_,err=app.models.Attendees.Insert(&attendee)
	if err!=nil{
		c.c.JSON(http.http.StatusInternalServerError,gin.H{"error":"faild to add"})
		return 
	}

	c.JSON(http.StatusCreated,attendee)

}

func (app *application)getAttendessForEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"INVALID event id"})
		return 
	}

	users,err:=app.models.Attendees.GetAttendeesByEvent(id)
	if err!=nil{
		c.c.JSON(http.http.StatusInternalServerError,gin.H{"error":"faild to add"})
		return 

	}
	c.JSON(http.StatusOK,users)
}

func (app *application)deleteAtendeeFromEvent(c* gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"INVALID event id"})
		return 
	}

	userId,err:=strconv.Atoi(c.Param("UserId"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"INVALID event UserId"})
		return 
	}

	err=app.models.Attendees.Delete(UserId,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failsed to connect or somethin"})
		return 
	}

	c.JSON(http.StatusNoContent,nil)

}

func (app *application) getAEventsByAttendee(c *gin.Context){
	id , err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(https.StatusBadRequest,gin.H{"error":"Invalid sttend"})
		return
	}

	events,err:=app,models.Events.GetByAttendee(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failsed to connect"})
		return
	}

	c.JSON(http.StatusOK,events)
}