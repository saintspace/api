package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	// Middleware to set CORS headers for preflight requests
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/email-subscriptions", postEmailSubscriptionsHandler)
		v1.POST("/email-verifications", postEmailVerificationsHandler)
	}
	return r
}
