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
		IsHTTPS:      true,
		ServerName:   "unsettled",
		Available:    false,
		SSLProtocols: "TLSv1 TLSv1.1 TLSv1.2",
		SSLCiphers:   "HIGH:!aNULL:!MD5",
		Rewrite:      "",
		Locations:    []*Location{&Location{From: "/", To: "http://127.0.0.1:8000/"}},
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
