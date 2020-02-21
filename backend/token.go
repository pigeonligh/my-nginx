package backend

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/my-nginx/utils"
)

// Token var
var Token string

// MakeToken function
func MakeToken() (string, int64) {
	timeout := time.Now().Add(time.Hour).Unix()
	key := utils.RandString(6)
	token := key + Token + strconv.FormatInt(timeout, 10)
	result := key + utils.MD5(token)

	return result, timeout
}

// CheckToken function
func CheckToken(target string, timeout int64) bool {
	key := target[0:6]
	token := key + Token + strconv.FormatInt(timeout, 10)
	result := key + utils.MD5(token)

	return result == target
}

// CheckLogged function
func CheckLogged(c *gin.Context) bool {
	token, err := c.Cookie("token")
	if err != nil {
		return false
	}
	stimeout, err := c.Cookie("timeout")
	if err != nil {
		return false
	}
	timeout, err := strconv.ParseInt(stimeout, 10, 64)
	if err != nil {
		return false
	}
	return CheckToken(token, timeout)
}

// Login Function
func Login(c *gin.Context) bool {
	password := c.DefaultQuery("password", "")
	if password == Token {
		token, timeout := MakeToken()
		c.SetCookie("token", token, 3600, "/", "", http.SameSiteDefaultMode, false, false)
		c.SetCookie("timeout", strconv.FormatInt(timeout, 10), 3600, "/", "", http.SameSiteDefaultMode, false, false)
		return true
	}
	return false
}

// Logout function
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", http.SameSiteDefaultMode, false, false)
	c.SetCookie("timeout", "", -1, "/", "", http.SameSiteDefaultMode, false, false)
}

func setupToken(r *gin.RouterGroup) {

	r.GET("check", func(c *gin.Context) {
		if CheckLogged(c) {
			utils.Response(c, 1, "", nil)
		} else {
			utils.Response(c, 0, "", nil)
		}
	})

}
