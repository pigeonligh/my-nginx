package ui

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/backend"
	"github.com/pigeonligh/my-nginx/httpconfig"
	"github.com/pigeonligh/my-nginx/utils"
)

func toString(locations []*httpconfig.Location) string {
	/*
		text := ""
		for _, location := range locations {
			text = text + location.From + " => " + location.To + "\n"
		}
	*/
	text, _ := json.MarshalIndent(locations, "", "	")
	return string(text)
}

func httpPage(c *gin.Context) {
	logged := backend.CheckLogged(c)
	if !logged {
		utils.Redirect(c, "./login")
		return
	}

	list := make([]gin.H, 0)
	indexes := make([]int, 0)

	for index := range backend.Data.HTTP.Data {
		indexes = append(indexes, index)
	}
	sort.Ints(indexes)

	for _, index := range indexes {
		config := backend.Data.HTTP.Data[index]
		name := "http://"
		if config.IsHTTPS {
			name = "https://"
		}
		obj := gin.H{
			"index":       config.Index,
			"name":        name + config.ServerName,
			"server_name": config.ServerName,
			"protocols":   config.SSLProtocols,
			"ciphers":     config.SSLCiphers,
			"rewrite":     config.Rewrite,
			"locations":   toString(config.Locations),
		}
		if config.IsHTTPS {
			obj["is_https"] = true
		} else {
			obj["is_http"] = true
		}
		if config.HTTPAttach == 0 {
			obj["http_none"] = true
		} else if config.HTTPAttach == 1 {
			obj["http_rewrite"] = true
		} else {
			obj["http_copy"] = true
		}
		if config.Available {
			obj["available"] = true
		}
		list = append(list, obj)
	}

	c.HTML(http.StatusOK, "http.html", gin.H{
		"title":  "HTTP 转发",
		"logged": logged,
		"list":   list,
	})
}

func httpModify(c *gin.Context) {
	logged := backend.CheckLogged(c)
	if !logged {
		utils.Redirect(c, "./login")
		return
	}

	index, err := strconv.Atoi(c.Query("i"))
	if err != nil {
		utils.Error404(c)
		return
	}
	config := backend.Data.HTTP.Data[index]
	if config == nil {
		utils.Error404(c)
		return
	}

	name := "http://"
	if config.IsHTTPS {
		name = "https://"
	}
	obj := gin.H{
		"index":       config.Index,
		"name":        name + config.ServerName,
		"server_name": config.ServerName,
		"protocols":   config.SSLProtocols,
		"ciphers":     config.SSLCiphers,
		"rewrite":     config.Rewrite,
		"locations":   toString(config.Locations),
	}
	if config.IsHTTPS {
		obj["is_https"] = true
	} else {
		obj["is_http"] = true
	}
	if config.HTTPAttach == 0 {
		obj["http_none"] = true
	} else if config.HTTPAttach == 1 {
		obj["http_rewrite"] = true
	} else {
		obj["http_copy"] = true
	}
	if config.Available {
		obj["available"] = true
	}

	c.HTML(http.StatusOK, "http_modify.html", gin.H{
		"title":  "HTTP 转发",
		"logged": logged,
		"config": obj,
	})
}
