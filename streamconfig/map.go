package streamconfig

// Map struct
type Map struct {
	MaxIndex int             `json:"max_index"`
	Data     map[int]*Config `json:"data"`
}

// New function
func (m *Map) New() {
	m.MaxIndex++
	m.Data[m.MaxIndex] = &Config{
		Index:          m.MaxIndex,
		Port:           65535,
		Dest:           "127.0.0.1:80",
		ConnectTimeout: 10,
		ProxyTimeout:   30,
	}
}

// NewMap function
func NewMap() *Map {
	return &Map{
		MaxIndex: 0,
		Data:     make(map[int]*Config),
	}
}
