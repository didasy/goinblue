package goblue

type SMS struct {
	To     string `json:"to"`
	From   string `json:"from"`
	Text   string `json:"text"`
	WebUrl string `json:"web_url"`
	Tag    string `json:"tag"`
	Type   string `json:"type"`
}
