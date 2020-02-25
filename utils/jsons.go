package utils

import (
	"github.com/gin-gonic/gin"
)

// HAccessDenied function
func HAccessDenied() gin.H {
	return gin.H{
		"status": 0,
		"msg":    "access denied",
		"data":   nil,
	}
}

// HError function
func HError(err error) gin.H {
	return gin.H{
		"status": 0,
		"msg":    err.Error(),
		"data":   nil,
	}
}

// HWarning function
func HWarning(msg string) gin.H {
	return gin.H{
		"status": 0,
		"msg":    msg,
		"data":   nil,
	}
}

// HResponse function
func HResponse(msg string, data interface{}) gin.H {
	return gin.H{
		"status": 1,
		"msg":    msg,
		"data":   data,
	}
}
