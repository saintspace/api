package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type PostVerifyEmailRequestData struct {
	SubscriptionToken string `json:"token"`
}

func (s *RouteHandler) PostVerifyEmailHandler(c *gin.Context) {
	var data PostVerifyEmailRequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error(fmt.Sprintf("error while parsing PostVerifyEmailHandler request => %v", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unescapedSubscriptionToken, err := url.QueryUnescape(data.SubscriptionToken)
	if err != nil {
		log.Error(fmt.Sprintf("error while unescaping email verification token => %v", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while parsing verification token"})
		return
	}
	err = s.emailService.VerifyEmailwithSubscriptionToken(unescapedSubscriptionToken)
	if err != nil {
		log.Error(fmt.Sprintf("error while verifying email {token: %s} => %s", unescapedSubscriptionToken, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error while attempting to verify subscription",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "email verification successful",
	})
}
