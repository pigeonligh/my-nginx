package nginx

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pigeonligh/my-nginx/utils"
)

// Run function
func Run() error {
	if err := utils.CheckPath("/etc/nginx/http.conf.d/"); err != nil {
		return err
	}
	if err := utils.CheckPath("/etc/nginx/stream.conf.d/"); err != nil {
		return err
	}
	if err := utils.CheckPath("/etc/nginx/certs/"); err != nil {
		return err
	}

	/*
		dst, err := os.OpenFile("/etc/nginx/nginx.conf", os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer dst.Close()
	*/
	src, err := os.Open("nginx/nginx.conf")
	if err != nil {
		return err
	}
	defer src.Close()
	bytes, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("/etc/nginx/nginx.conf", bytes, 0666)
	if err != nil {
		return err
	}

	cmd := exec.Command("sh", "-c", "nginx")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Reload function
func Reload() error {
	cmd := exec.Command("sh", "-c", "nginx -s reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
