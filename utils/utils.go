package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// CheckPath function
func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		fmt.Println("create dir " + path)
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

func init() {
	rand.Seed(time.Now().Unix())
}
