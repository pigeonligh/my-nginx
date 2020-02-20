package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/backend"
	"github.com/pigeonligh/my-nginx/utils"
)

func sslPage(c *gin.Context) {
	logged := backend.CheckLogged(c)
	if !logged {
		utils.Redirect(c, "./login")
		return
	}

	c.HTML(http.StatusOK, "ssl.html", gin.H{
		"title":  "SSL 证书管理",
		"logged": logged,
	})
}
