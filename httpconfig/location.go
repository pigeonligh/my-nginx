package httpconfig

import "fmt"

// Location struct
type Location struct {
	From      string `json:"from"`
	To        string `json:"to"`
	WebSocket bool   `json:"websocket" defalut:"true"`
}

// WriteString function
func (c Location) WriteString() string {
	var data string = ""
	data = fmt.Sprintf("%s\n location %s {", data, c.From)
	data = fmt.Sprintf("%s\n proxy_pass %s;", data, c.To)
	data = fmt.Sprintf("%s\n proxy_set_header Host $host;", data)
	data = fmt.Sprintf("%s\n proxy_set_header X-Real-IP $remote_addr;", data)
	data = fmt.Sprintf("%s\n proxy_set_header REMOTE-HOST $remote_addr;", data)
	data = fmt.Sprintf("%s\n proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;", data)
	data = fmt.Sprintf("%s\n proxy_set_header Via \"my-nginx\";", data)
	if c.WebSocket {
		data = fmt.Sprintf("%s\n proxy_http_version 1.1;", data)
		data = fmt.Sprintf("%s\n proxy_set_header Upgrade $http_upgrade;", data)
		data = fmt.Sprintf("%s\n proxy_set_header Connection \"upgrade\";", data)
	}
	data = fmt.Sprintf("%s\n }\n", data)
	return data
}
