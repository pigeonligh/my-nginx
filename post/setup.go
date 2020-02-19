package post

import (
	"github.com/gin-gonic/gin"
)

// Setup function
func Setup(r *gin.RouterGroup) {
	setupToken(r.Group("sign"))
}
