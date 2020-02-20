package backend

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pigeonligh/my-nginx/httpconfig"
	"github.com/pigeonligh/my-nginx/nginx"
	"github.com/pigeonligh/my-nginx/ssl"
	"github.com/pigeonligh/my-nginx/streamconfig"
)

// Configs struct
type Configs struct {
	NewModify bool              `json:"changed"`
	HTTP      *httpconfig.Map   `json:"http"`
	Stream    *streamconfig.Map `json:"stream"`
	SSL       *ssl.Map          `json:"ssl"`
}

// Data var
var Data *Configs

// Save function
func (c *Configs) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("/etc/nginx/my-nginx.json", data, 0777)
}

// NewConfigs function
func NewConfigs() *Configs {
	configs := &Configs{
		HTTP:   httpconfig.NewMap(),
		Stream: streamconfig.NewMap(),
		SSL:    ssl.NewMap(),
	}
	configs.Save()
	return configs
}

// LoadConfigs function
func LoadConfigs() (*Configs, error) {
	_, err := os.Stat("/etc/nginx/my-nginx.json")
	if os.IsNotExist(err) {
		return NewConfigs(), nil
	}
	file, err := os.Open("/etc/nginx/my-nginx.json")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var configs *Configs
	err = json.Unmarshal(data, configs)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

// Apply function
func (c *Configs) Apply() error {
	if !c.NewModify {
		return nil
	}
	var err error
	err = c.HTTP.Apply(c.SSL)
	if err != nil {
		return err
	}
	err = c.Stream.Apply()
	if err != nil {
		return err
	}
	err = nginx.Reload()
	if err != nil {
		return err
	}
	c.NewModify = false
	return nil
}
