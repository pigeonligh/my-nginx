package streamconfig

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// Config struct
type Config struct {
	Index          int    `json:"index"`
	Port           int    `json:"port"`
	Dest           string `json:"dest"`
	ConnectTimeout int    `json:"connect_timeout"`
	ProxyTimeout   int    `json:"timeout"`
}

// WriteConfig function
func (c Config) WriteConfig() error {
	filename := "/etc/nginx/stream.conf.d/my-nginx." + strconv.Itoa(c.Index) + ".conf"

	var data string = ""
	data = fmt.Sprintf("%s\nserver {", data)
	data = fmt.Sprintf("%s\n listen %d;", data, c.Port)
	data = fmt.Sprintf("%s\n proxy_pass %s;", data, c.Dest)
	data = fmt.Sprintf("%s\n proxy_timeout %ds;", data, c.ProxyTimeout)
	data = fmt.Sprintf("%s\n proxy_connect_timeout %ds;", data, c.ConnectTimeout)
	data = fmt.Sprintf("%s\n}\n", data)

	return ioutil.WriteFile(filename, []byte(data), 0666)
}
