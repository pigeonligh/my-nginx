package ui

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/backend"
	"github.com/pigeonligh/my-nginx/utils"
)

func portPage(c *gin.Context) {
	logged := backend.CheckLogged(c)
	if !logged {
		utils.Redirect(c, "./login")
		return
	}

	list := make([]gin.H, 0)
	for _, config := range backend.Data.Stream.Data {
		list = append(list, gin.H{
			"index": config.Index,
			"name":  strconv.Itoa(config.Port) + " => " + config.Dest,
		})
	}

	c.HTML(http.StatusOK, "port.html", gin.H{
		"title":  "端口转发",
		"logged": logged,
		"list":   list,
	})
}
