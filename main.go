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

	ui.Setup(r.Group(""))
	backend.Setup(r.Group("/apis"))

	/*
		r.GET("", func(c *gin.Context) {
			str := "server {\n" +
				"listen       80 default_server;\n" +
				"server_name  _;\n" +
				"location / {\n" +
				"deny all;\n" +
				"}\n}"

			path := "/etc/nginx/http.conf.d/test.conf"

			if err := ioutil.WriteFile(path, []byte(str), 0777); err != nil {
				c.String(http.StatusBadRequest, "bad\n"+err.Error())
				return
			}
			if err := nginx.Reload(); err != nil {
				c.String(http.StatusBadRequest, "bad\n"+err.Error())
				return
			}
			c.String(http.StatusOK, "test "+token)

		})
	*/
	r.Run(":8080")
}
