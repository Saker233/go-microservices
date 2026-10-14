package main

import (
	"log-service/cmd/api/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) writeLog(c *gin.Context) {
	// read json into var
	var req JSONPayload
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event := data.LogEntry{
		Name: req.Name,
		Data: req.Data,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "logged",
	}

	c.JSON(http.StatusAccepted, resp)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
