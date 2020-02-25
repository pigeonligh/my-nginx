package backend

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func getSSL(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		return utils.HError(err)
	}
	config := Data.SSL.Data[index]
	if config == nil {
		return utils.HWarning("config not exists")
	}

	return utils.HResponse("", config)
}

func getHTTP(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		return utils.HError(err)
	}
	config := Data.HTTP.Data[index]
	if config == nil {
		return utils.HWarning("config not exists")
	}

	return utils.HResponse("", config)
}

func setupGET(r *gin.RouterGroup) {
	r.GET("ssl", func(c *gin.Context) {
		c.JSON(http.StatusOK, getSSL(c))
	})
	r.GET("http", func(c *gin.Context) {
		c.JSON(http.StatusOK, getHTTP(c))
	})
}
