package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

// Setup function
func Setup(r *gin.RouterGroup) {
	setupToken(r.Group("sign"))
	setupAdd(r.Group("add"))
	setupModify(r.Group("modify"))
	setupGET(r.Group("get"))
	r.GET("reload", func(c *gin.Context) {
		if err := Data.Apply(); err != nil {
			utils.Response(c, 0, err.Error(), nil)
		}
		utils.Response(c, 1, "", nil)
	})
}
