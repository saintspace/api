package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler iHandler
}

func New(handler iHandler) *Router {
	return &Router{
		handler: handler,
	}
}

type iHandler interface {
	PostEmailSubscriptionsHandler(c *gin.Context)
	PostVerifyEmailHandler(c *gin.Context)
}

func (s *Router) GetRouter() *gin.Engine {
	r := gin.Default()
	// Middleware to set CORS headers for preflight requests
	r.Use(func(c *gin.Context) {
		log.Println("in CORS middlware")
		if c.Request.Method == "OPTIONS" {
			log.Println("is OPTIONS request")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		log.Println("not OPTIONS request")
		c.Next()
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/email-subscriptions", s.handler.PostEmailSubscriptionsHandler)
		v1.POST("/email-verifications", s.handler.PostVerifyEmailHandler)
	}
	return r
}
