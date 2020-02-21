package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/backend"
	"github.com/pigeonligh/my-nginx/nginx"
	"github.com/pigeonligh/my-nginx/ui"
)

func main() {
	flag.StringVar(&backend.Token, "token", "", "control token")
	flag.Parse()

	if err := nginx.Run(); err != nil {
		log.Fatal(err)
		return
	}

	data, err := backend.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}
	backend.Data = data

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.StaticFile("favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("templates/**")

	ui.Setup(r.Group(""))
	backend.Setup(r.Group("/apis"))
	r.Run(":8080")
}
