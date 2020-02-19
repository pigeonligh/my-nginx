package ssl

import "io"

// Map struct
type Map struct {
	MaxIndex int             `json:"max_index"`
	Data     map[int]*Config `json:"data"`
}

// New function
func (m *Map) New(domain string, crt io.Reader, key io.Reader) error {
	m.MaxIndex++
	m.Data[m.MaxIndex] = &Config{
		Index:      m.MaxIndex,
		DomainName: domain,
	}
	return m.Data[m.MaxIndex].Save(crt, key)
}

func check(str, domain string) bool {
	if str[0] == '*' {
		if len(str) == 1 {
			return true
		}
		if str[1] == '.' && str[2:len(str)] == domain {
			return true
		}
		prefixLen := len(domain) - len(str) + 1
		if prefixLen <= 0 {
			return false
		}
		prefix := domain[0:prefixLen]
		str = prefix + str[1:len(str)]
		return str == domain
	}
	return str == domain
}

// Find function
func (m *Map) Find(domain string) *Config {
	for _, config := range m.Data {
		if check(config.DomainName, domain) {
			return config
		}
	}
	return nil
}

// NewMap function
func NewMap() *Map {
	return &Map{
		MaxIndex: 0,
		Data:     make(map[int]*Config),
	}
}
