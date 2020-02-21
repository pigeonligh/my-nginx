package backend

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/httpconfig"
	"github.com/pigeonligh/my-nginx/utils"
)

func modifySSL(c *gin.Context) {
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

	crt := c.PostForm("crt")
	key := c.PostForm("key")
	domain := c.PostForm("domain")

	config.DomainName = domain
	Data.NewModify = true
	if crt != "" && key != "" {
		err = config.Save([]byte(crt), []byte(key))
		if err != nil {
			utils.Response(c, 0, err.Error(), nil)
			return
		}
	}
	if err = Data.Save(); err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	utils.Redirect(c, "./../../saved")
}

func modifyHTTP(c *gin.Context) {
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

	var isHTTPS bool
	var serverName string
	var available bool
	var protocols string
	var ciphers string
	var rewrite string
	var locations []*httpconfig.Location

	isHTTPS = bool(c.PostForm("is_https") == "true")
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
			utils.Response(c, 0, err.Error(), nil)
			return
		}
	} else {
		locations = make([]*httpconfig.Location, 0)
	}

	config.IsHTTPS = isHTTPS
	config.Available = available
	config.ServerName = serverName
	config.SSLProtocols = protocols
	config.SSLCiphers = ciphers
	config.Rewrite = rewrite
	config.Locations = locations
	Data.NewModify = true
	if err = Data.Save(); err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	utils.Redirect(c, "./../../saved")
}

func modifyStream(c *gin.Context) {
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

	var port int64
	var dest string
	var connectTimeout int64
	var proxyTimeout int64

	port, err = strconv.ParseInt(c.PostForm("port"), 10, 32)
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	connectTimeout, err = strconv.ParseInt(c.PostForm("connect_timeout"), 10, 32)
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	proxyTimeout, err = strconv.ParseInt(c.PostForm("timeout"), 10, 32)
	if err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	dest = c.PostForm("dest")

	config.Port = int(port)
	config.Dest = dest
	config.ConnectTimeout = int(connectTimeout)
	config.ProxyTimeout = int(proxyTimeout)

	Data.NewModify = true
	if err = Data.Save(); err != nil {
		utils.Response(c, 0, err.Error(), nil)
		return
	}
	utils.Redirect(c, "./../../saved")
}

func setupModify(r *gin.RouterGroup) {
	r.POST("ssl", modifySSL)
	r.POST("http", modifyHTTP)
	r.POST("stream", modifyStream)
}
