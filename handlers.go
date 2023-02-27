package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostEmailSubscriptionRequestData struct {
	EmailAddress string `json:"email"`
}

var postEmailSubscriptionsHandler = func(c *gin.Context) {

	var data PostEmailSubscriptionRequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !emailIsValid(data.EmailAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email address"})
		return
	}
	if err := createEmailSubscription(data.EmailAddress); err != nil {
		log.Printf("error while creating email subscription => %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while saving your email subscription"})
		return
	}

	if err := publishTask("test task from api"); err != nil {
		log.Printf("error while publishing email verification task => %v", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "email subscription successful",
		"email":   data.EmailAddress,
	})
}

var postEmailVerificationsHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "email verification successful",
	})
}
