package main

import (
	"encoding/json"
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

	emailSendTask := EmailSendTask{
		TemplateName:  "email-subscription-verification",
		SenderAddress: "noreply@dev-messages.saintspace.app",
		SubjectLine:   "SaintSpace: Confirm Your Subscription",
		ToAddresses:   []string{data.EmailAddress},
		Parameters: EmailSendTaskParameters{
			VerificationLink: "https://dev.saintspace.app/p/verify-email?token=",
		},
	}
	emailSendTaskBytes, err := json.Marshal(emailSendTask)
	if err != nil {
		log.Printf("error while marshaling email send task details => %v", err.Error())
	} else {
		task := Task{
			TaskName:      "email-send",
			CorrelationId: "",
			TaskDetails:   string(emailSendTaskBytes),
		}
		taskBytes, err := json.Marshal(task)
		if err != nil {
			log.Printf("error while marshaling email send task => %v", err.Error())
		} else {
			if err := publishTask(string(taskBytes)); err != nil {
				log.Printf("error while publishing email send task => %v", err.Error())
			}
		}
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
