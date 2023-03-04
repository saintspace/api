package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostEmailSubscriptionRequestData struct {
	EmailAddress string `json:"email"`
}

func (s *RouteHandler) PostEmailSubscriptionsHandler(c *gin.Context) {
	var data PostEmailSubscriptionRequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !s.emailService.IsValidEmail(data.EmailAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email address"})
		return
	}
	subscriptionExists, err := s.emailService.EmailSubscriptionExists(data.EmailAddress)
	if err != nil {
		log.Printf("error while checking if email subscription already exists => %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while processing your email subscription"})
		return
	}

	if !subscriptionExists {
		if err := s.emailService.CreateEmailSubscription(data.EmailAddress); err != nil {
			log.Printf("error while creating email subscription => %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while saving your email subscription"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "email subscription successful",
		"email":   data.EmailAddress,
	})
}
