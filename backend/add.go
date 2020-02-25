package backend

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func newSSL(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	crt := []byte(c.PostForm("crt"))
	key := []byte(c.PostForm("key"))
	domain := c.PostForm("domain")

	index, err := Data.SSL.New(domain, crt, key)
	if err != nil {
		return utils.HError(err)
	}
	if err = Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse(strconv.Itoa(index), nil)
}

func newHTTP(c *gin.Context) gin.H {
	if !CheckLogged(c) {
		return utils.HAccessDenied()
	}

	index := Data.HTTP.New()
	if err := Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse(strconv.Itoa(index), nil)
}

func setupAdd(r *gin.RouterGroup) {
	r.POST("ssl", func(c *gin.Context) {
		c.JSON(http.StatusOK, newSSL(c))
	})
	r.POST("http", func(c *gin.Context) {
		c.JSON(http.StatusOK, newHTTP(c))
	})
}
