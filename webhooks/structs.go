package webhooks

import "time"

type Webhook struct {
	ID    string    `json:"id"`
	Event string    `json:"event"`
	Url   string    `json:"url"`
	Time  time.Time `json:"time"`
}

type Invocation struct {
	Event      string `json:"event"`
	Parameters string `json:"parameters"`
	Timestamp  string `json:"time"`
}

var Wh = Webhook{}
