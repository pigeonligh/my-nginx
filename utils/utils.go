package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// CheckPath function
func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.Mkdir(path, 0777)
	}
	return err
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// RandString function
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// MD5 function
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// Response function
func Response(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}

// Error404 function
func Error404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"title": "404 Not Found",
	})
}

// Redirect function
func Redirect(c *gin.Context, dest string) {
	c.Header("Cache-Control", "must-revalidate, no-store")
	c.Header("Location", dest)
	c.String(http.StatusTemporaryRedirect, "")
}

func init() {
	rand.Seed(time.Now().Unix())
}
