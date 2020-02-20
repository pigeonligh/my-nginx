package backend

import (
	"github.com/gin-gonic/gin"
)

// Setup function
func Setup(r *gin.RouterGroup) {
	setupToken(r.Group("sign"))
	setupAdd(r.Group("add"))
	setupModify(r.Group("modify"))
	setupGET(r.Group("get"))
}
