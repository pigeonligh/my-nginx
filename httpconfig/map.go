package httpconfig

import "github.com/pigeonligh/my-nginx/ssl"

// Map struct
type Map struct {
	MaxIndex int             `json:"max_index"`
	Data     map[int]*Config `json:"data"`
}

// New function
func (m *Map) New() int {
	m.MaxIndex++
	m.Data[m.MaxIndex] = &Config{
		Index:        m.MaxIndex,
		IsHTTPS:      false,
		ServerName:   "unsettled",
		Available:    false,
		SSLProtocols: "",
		SSLCiphers:   "",
		Rewrite:      "",
		Locations:    make([]*Location, 0),
	}
	return m.MaxIndex
}

// NewMap function
func NewMap() *Map {
	return &Map{
		MaxIndex: 0,
		Data:     make(map[int]*Config),
	}
}

// Apply function
func (m *Map) Apply(sslMap *ssl.Map) error {
	for _, config := range m.Data {
		err := config.WriteConfig(sslMap)
		if err != nil {
			return err
		}
	}
	return nil
}
