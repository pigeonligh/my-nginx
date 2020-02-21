package ui

import (
	"net/http"
	"sort"

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

	list := make([]gin.H, 0)
	indexes := make([]int, 0)

	for index := range backend.Data.SSL.Data {
		indexes = append(indexes, index)
	}
	sort.Ints(indexes)

	for _, index := range indexes {
		config := backend.Data.SSL.Data[index]
		obj := gin.H{
			"index":  config.Index,
			"domain": config.DomainName,
		}
		list = append(list, obj)
	}

	c.HTML(http.StatusOK, "ssl.html", gin.H{
		"title":  "SSL 证书管理",
		"logged": logged,
		"list":   list,
	})
}
