package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/httpconfig"
	"github.com/pigeonligh/my-nginx/utils"
)

func modifySSL(c *gin.Context) gin.H {
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

	crt := c.PostForm("crt")
	key := c.PostForm("key")
	domain := c.PostForm("domain")

	config.DomainName = domain
	if crt != "" && key != "" {
		err = config.Save([]byte(crt), []byte(key))
		if err != nil {
			return utils.HError(err)
		}
	}
	if err = Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse("", nil)
}

func modifyHTTP(c *gin.Context) gin.H {
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

	var isHTTPS bool
	var httpAttach int
	var serverName string
	var available bool
	var protocols string
	var ciphers string
	var rewrite string
	var locations []*httpconfig.Location

	isHTTPS = bool(c.PostForm("is_https") == "true")
	httpAttach, err = strconv.Atoi(c.PostForm("http_attach"))
	if err != nil {
		return utils.HError(err)
	}
	available = bool(c.PostForm("available") == "true")
	fmt.Println(isHTTPS, available)
	serverName = c.PostForm("server_name")
	rewrite = c.PostForm("rewrite")

	if isHTTPS {
		protocols = c.PostForm("ssl_protocols")
		ciphers = c.PostForm("ssl_ciphers")
	} else {
		protocols = ""
		ciphers = ""
	}
	if rewrite == "" {
		locationString := c.PostForm("locations")
		err = json.Unmarshal([]byte(locationString), &locations)
		if err != nil {
			return utils.HError(err)
		}
	} else {
		locations = make([]*httpconfig.Location, 0)
	}

	config.IsHTTPS = isHTTPS
	config.HTTPAttach = httpAttach
	config.Available = available
	config.ServerName = serverName
	config.SSLProtocols = protocols
	config.SSLCiphers = ciphers
	config.Rewrite = rewrite
	config.Locations = locations
	if err = Data.Save(); err != nil {
		return utils.HError(err)
	}
	return utils.HResponse("", nil)
}

func setupModify(r *gin.RouterGroup) {
	r.POST("ssl", func(c *gin.Context) {
		c.JSON(http.StatusOK, modifySSL(c))
	})
	r.POST("http", func(c *gin.Context) {
		c.JSON(http.StatusOK, modifyHTTP(c))
	})
}
