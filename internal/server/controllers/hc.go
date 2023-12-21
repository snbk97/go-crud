package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"date":    time.Now().Format(time.ANSIC),
	})
}
