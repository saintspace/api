package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostVerifyEmailRequestData struct {
	EmailAddress      string `json:"email"`
	SubscriptionToken string `json:"token"`
}

func (s *RouteHandler) PostVerifyEmailHandler(c *gin.Context) {
	var data PostVerifyEmailRequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "email verification successful",
		"email":   data.EmailAddress,
	})
}
