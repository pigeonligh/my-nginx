package backend

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func delSSL(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		return utils.HError(err)
	}
	config := Data.SSL.Data[index]
	if config == nil {
		return utils.HWarning("config not exists")
	}

	delete(Data.SSL.Data, index)
	if err = Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse("", nil)
}

func delHTTP(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		return utils.HError(err)
	}
	config := Data.HTTP.Data[index]
	if config == nil {
		return utils.HWarning("config not exists")
	}

	delete(Data.HTTP.Data, index)
	if err = Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse("", nil)
}

func setupDEL(r *gin.RouterGroup) {
	r.POST("ssl", func(c *gin.Context) {
		c.JSON(http.StatusOK, delSSL(c))
	})
	r.POST("http", func(c *gin.Context) {
		c.JSON(http.StatusOK, delHTTP(c))
	})
}
