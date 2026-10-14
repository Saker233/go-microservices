package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *Config) authenticate(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// valiudate user against db
	user, err := app.Models.User.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	valid, err := user.PasswordMatches(req.Password)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	c.JSON(http.StatusAccepted, payload)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
