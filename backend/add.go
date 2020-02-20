package backend

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func newSSL(c *gin.Context) {
	if !checkLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	crt := []byte(c.PostForm("crt"))
	key := []byte(c.PostForm("key"))
	domain := c.PostForm("domain")

	index, err := Data.SSL.New(domain, crt, key)
	Data.NewModify = true
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	utils.Response(c, 1, strconv.Itoa(index), nil)
}

func newHTTP(c *gin.Context) {
	if !checkLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index := Data.HTTP.New()
	Data.NewModify = true
	utils.Response(c, 1, strconv.Itoa(index), nil)
}

func newStream(c *gin.Context) {
	if !checkLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index := Data.Stream.New()
	Data.NewModify = true
	utils.Response(c, 1, strconv.Itoa(index), nil)
}

func setupAdd(r *gin.RouterGroup) {
	r.POST("ssl", newSSL)
	r.POST("http", newHTTP)
	r.POST("stream", newStream)
}
