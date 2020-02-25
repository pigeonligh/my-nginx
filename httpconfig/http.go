package httpconfig

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/pigeonligh/my-nginx/ssl"
)

// Config struct
type Config struct {
	Index        int         `json:"index"`
	IsHTTPS      bool        `json:"is_https" default:"true"`
	HTTPAttach   int         `json:"http_attach" defalut:"0"`
	ServerName   string      `json:"server_name"`
	Available    bool        `json:"available" defalut:"false"`
	SSLProtocols string      `json:"ssl_protocols"`
	SSLCiphers   string      `json:"ssl_ciphers"`
	Rewrite      string      `json:"rewrite"`
	Locations    []*Location `json:"locations"`
}

// WriteConfig function
func (c Config) WriteConfig(sslMap *ssl.Map) error {
	filename := "/etc/nginx/http.conf.d/my-nginx." + strconv.Itoa(c.Index) + ".conf"
	defaultString := ""
	if c.ServerName == "_" {
		defaultString = " default"
	}

	var data string = ""
	var body string = ""
	data = fmt.Sprintf("%s\nserver {", data)
	if c.IsHTTPS {
		data = fmt.Sprintf("%s\n listen 443 ssl%s;", data, defaultString)

		sslConfig := sslMap.Find(c.ServerName)
		if sslConfig != nil {
			data = fmt.Sprintf("%s\n ssl_certificate %s;", data, sslConfig.GetCrtPath())
			data = fmt.Sprintf("%s\n ssl_certificate_key %s;", data, sslConfig.GetKeyPath())
			if c.SSLProtocols != "" {
				data = fmt.Sprintf("%s\n ssl_protocols %s;", data, c.SSLProtocols)
			}
			if c.SSLCiphers != "" {
				data = fmt.Sprintf("%s\n ssl_ciphers %s;", data, c.SSLCiphers)
			}
		}
	} else {
		data = fmt.Sprintf("%s\n listen 80%s;", data, defaultString)
	}
	data = fmt.Sprintf("%s\n server_name %s;", data, c.ServerName)
	if c.Available {
		if c.Rewrite != "" {
			body = fmt.Sprintf("%s\n rewrite ^/(.*)$ %s/$1 permanent;", body, c.Rewrite)
		} else {
			for _, location := range c.Locations {
				body = fmt.Sprintf("%s\n %s", body, location.WriteString())
			}
		}
	} else {
		body = fmt.Sprintf("%s\n location / {\n  deny all;\n }", body)
	}
	data = fmt.Sprintf("%s\n%s", data, body)
	data = fmt.Sprintf("%s\n}\n", data)

	if c.IsHTTPS && c.HTTPAttach > 0 {
		data = fmt.Sprintf("%s\nserver {", data)
		data = fmt.Sprintf("%s\n listen 80%s;", data, defaultString)
		data = fmt.Sprintf("%s\n server_name %s;", data, c.ServerName)
		if c.Available {
			if c.HTTPAttach == 1 {
				data = fmt.Sprintf("%s\n rewrite ^/(.*)$ https://%s/$1 permanent;", data, c.ServerName)
			} else {
				data = fmt.Sprintf("%s\n %s", data, body)
			}
		} else {
			data = fmt.Sprintf("%s\n location / {\n  deny all;\n }", data)
		}
		data = fmt.Sprintf("%s\n}\n", data)
	}

	return ioutil.WriteFile(filename, []byte(data), 0666)
}

// NewConfig function
func NewConfig(index int) *Config {
	return &Config{
		Index:        index,
		IsHTTPS:      true,
		HTTPAttach:   0,
		ServerName:   "unsettled",
		Available:    false,
		SSLProtocols: "TLSv1 TLSv1.1 TLSv1.2",
		SSLCiphers:   "HIGH:!aNULL:!MD5",
		Rewrite:      "",
		Locations:    []*Location{&Location{From: "/", To: "http://127.0.0.1:8000/", WebSocket: true}},
	}
}
