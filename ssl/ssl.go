package ssl

import (
	"io/ioutil"
	"strconv"
)

// Config struct
type Config struct {
	Index      int    `json:"index"`
	DomainName string `json:"domain"`
}

// Save function
func (c Config) Save(crt []byte, key []byte) error {
	var err error
	err = ioutil.WriteFile(c.GetCrtPath(), crt, 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.GetKeyPath(), key, 0777)
	if err != nil {
		return err
	}
	return nil
}

// GetDirPath function
func (c Config) GetDirPath() string {
	dir := "/etc/nginx/certs/" + strconv.Itoa(c.Index) + "."
	return dir
}

// GetCrtPath function
func (c Config) GetCrtPath() string {
	dir := c.GetDirPath()
	return dir + "ssl.crt"
}

// GetKeyPath function
func (c Config) GetKeyPath() string {
	dir := c.GetDirPath()
	return dir + "ssl.key"
}
