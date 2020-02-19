package httpconfig

// Map struct
type Map struct {
	MaxIndex int             `json:"max_index"`
	Data     map[int]*Config `json:"data"`
}

// New function
func (m *Map) New() {
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
}

// NewMap function
func NewMap() *Map {
	return &Map{
		MaxIndex: 0,
		Data:     make(map[int]*Config),
	}
}
