package ssl

import (
	"io"
	"os"
	"strconv"

	"github.com/pigeonligh/my-nginx/utils"
)

// Config struct
type Config struct {
	Index      int    `json:"index"`
	DomainName string `json:"domain"`
}

// Save function
func (c Config) Save(crt io.Reader, key io.Reader) error {
	if err := utils.CheckPath(c.GetDirPath()); err != nil {
		return err
	}
	crtFile, err := os.OpenFile(c.GetCrtPath(), os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer crtFile.Close()

	keyFile, err := os.OpenFile(c.GetKeyPath(), os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	if _, err := io.Copy(crtFile, crt); err != nil {
		return err
	}
	if _, err := io.Copy(keyFile, crt); err != nil {
		return err
	}
	return nil
}

// GetDirPath function
func (c Config) GetDirPath() string {
	dir := "/etc/nginx/certs/" + strconv.Itoa(c.Index) + "/"
	utils.CheckPath(dir)
	return dir
}

// GetCrtPath function
func (c Config) GetCrtPath() string {
	return c.GetDirPath() + "ssl.crt"
}

// GetKeyPath function
func (c Config) GetKeyPath() string {
	return c.GetKeyPath() + "ssl.key"
}
