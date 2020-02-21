package backend

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

func getSSL(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.SSL.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	utils.Response(c, 1, "", config)
}

func getHTTP(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.HTTP.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	utils.Response(c, 1, "", config)
}

func getStream(c *gin.Context) {
	if !CheckLogged(c) {
		utils.Response(c, 0, "access denied", nil)
		return
	}

	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	config := Data.Stream.Data[index]
	if config == nil {
		utils.Response(c, 0, "config not exists", nil)
		return
	}

	utils.Response(c, 1, "", config)
}

func setupGET(r *gin.RouterGroup) {
	r.GET("ssl", getSSL)
	r.GET("http", getHTTP)
	r.GET("stream", getStream)
}
