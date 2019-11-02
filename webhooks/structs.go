package webhooks

type Webhooks struct {
	Webhooks []Webhook
}

type Webhook struct {
	Id    string `json:"id"`
	Event string `json:"event"`
	Url   string `json:"url"`
	Time  string `json:"time"`
}

type Invocation struct {
	Event  string `json:"event"`
	Params string `json:"parameters"`
	Time   string `json:"time"`
}
