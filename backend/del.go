package backend

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func delSSL(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.SSL.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	delete(Data.SSL.Data, index)

	utils.Response(c, 1, "", nil)
}

func delHTTP(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.HTTP.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	delete(Data.HTTP.Data, index)

	utils.Response(c, 1, "", nil)
}

func delStream(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.PostForm("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.Stream.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	delete(Data.SSL.Data, index)

	utils.Response(c, 1, "", nil)
}

func setupDEL(r *gin.RouterGroup) {
	r.POST("ssl", delSSL)
	r.POST("http", delHTTP)
	r.POST("stream", delStream)
}
