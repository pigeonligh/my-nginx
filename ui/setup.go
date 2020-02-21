package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/backend"
	"github.com/pigeonligh/my-nginx/utils"
)

// Setup function
func Setup(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		logged := backend.CheckLogged(c)
		if !logged {
			utils.Redirect(c, "./login")
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":  "My Nginx",
			"logged": logged,
		})
	})

	r.GET("login", func(c *gin.Context) {
		logged := backend.CheckLogged(c)
		if logged {
			utils.Redirect(c, "./")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"logged": logged,
		})
	})

	r.GET("checklogin", func(c *gin.Context) {
		if backend.Login(c) {
			utils.Redirect(c, "./login")
		} else {
			utils.Redirect(c, "./login?type=error")
		}
	})

	r.GET("logout", func(c *gin.Context) {
		backend.Logout(c)
		utils.Redirect(c, "./login")
	})

	r.GET("apply", func(c *gin.Context) {
		logged := backend.CheckLogged(c)
		if !logged {
			utils.Redirect(c, "./login")
			return
		}

		if err := backend.Data.Apply(); err != nil {
			c.String(http.StatusConflict, "<pre>Error: \n"+err.Error())
			return
		}
		if err := backend.Data.Save(); err != nil {
			c.String(http.StatusConflict, "<pre>Error: \n"+err.Error())
			return
		}
		utils.Redirect(c, "./success")
	})

	r.GET("success", func(c *gin.Context) {
		logged := backend.CheckLogged(c)
		if !logged {
			utils.Redirect(c, "./login")
			return
		}

		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":  "My Nginx",
			"logged": logged,
		})
	})

	r.GET("http", httpPage)

	r.GET("ssl", sslPage)
}
