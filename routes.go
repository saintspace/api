package main

import (
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/email-subscriptions", postEmailSubscriptionsHandler)
		v1.POST("/email-verifications", postEmailVerificationsHandler)
	}
	return r
}
