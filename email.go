package goinblue

// Email request to be send
type Email struct {
	To          map[string]string `json:"to"`
	Subject     string            `json:"subject"`
	From        []string          `json:"from"`
	Html        string            `json:"html"`
	Text        string            `json:"text"`
	Cc          map[string]string `json:"cc"`
	Bcc         map[string]string `json:"bcc"`
	ReplyTo     []string          `json:"replyto"`
	Attachment  interface{}       `json:"attachment"`
	Headers     map[string]string `json:"headers"`
	InlineImage map[string]string `json:"inline_image"`
}

// Email template request to be send
type EmailTemplate struct {
	Id            int               `json:"id"`
	To            map[string]string `json:"to"`
	Cc            map[string]string `json:"cc"`
	Bcc           map[string]string `json:"bcc"`
	Attr          map[string]string `json:"attr"`
	AttachmentUrl []string          `json:"attachment_url"`
	Attachment    map[string]string `json:"attachment"`
	Headers       map[string]string `json:"headers"`
}

// This is here for documentation purpose
type Attachment map[string]string

// This is here for documentation purpose
type AttachmentUrl []string
