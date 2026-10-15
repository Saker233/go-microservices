package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type mailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) sendMail(c *gin.Context) {
	var req mailMessage

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	msg := Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Data:    req.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := jsonResponse {
		Error: false,
		Message: "sent to " + req.To,
	}

	c.JSON(http.StatusAccepted, payload)

}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
