package main

import (
	"github.com/gin-gonic/gin"
)

var postEmailSubscriptionsHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "email subscription successful",
	})
}

var postEmailVerificationsHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "email verification successful",
	})
}
