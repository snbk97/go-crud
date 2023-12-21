package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleOrmError(c *gin.Context, o *gorm.DB, message string) {
	fmt.Printf("HandleOrmError: %v\n", o)
	if o.Error != nil || o.RowsAffected == 0 {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		panic(message)
	}
}

func CreateHttpError(c *gin.Context, err error, statusCode int, message string) bool {
	if err == nil {
		return false
	}

	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	logger.LogErr(err, message)
	c.JSON(statusCode, gin.H{
		"error": message,
	})

	return true
}
