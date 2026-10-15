package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type mailMessage struct {
	From string `json:"from"`
	To 		string	`json:"to"`
	
}

func (app *Config) sendMail(c * gin.Context) {

}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
