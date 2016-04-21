// Golang library for sendinblue API

package goblue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BASE_URL            = "https://api.sendinblue.com/v2.0"
	CONTENT_TYPE        = "application/json"
	EMAIL_URL           = "/email"
	EMAIL_TEMPLATE_URL  = "/template"
	POST                = "POST"
	CONTENT_TYPE_HEADER = "Content-Type"
	API_KEY_HEADER      = "api-key"
)

var (
	DEFAULT_SEND_TIMEOUT time.Duration = time.Second * 10
)

// The main struct of this package
type Goblue struct {
	ApiKey            string
	Timeout           time.Duration
	BaseUrl           string
	ContentType       string
	Method            string
	ContentTypeHeader string
	ApiKeyHeader      string
	EmailUrl          string
	EmailTemplateUrl  string
}

// The response from server
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// To get message-id of a sent message
func (r *Response) GetMessageId() (string, error) {
	dataInterface, ok := r.Data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Invalid Data type: ", "Cannot convert to map[string]interface{}")
	}
	emailID, ok := dataInterface["message-id"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid Data type: ", "message-id is not a string")
	}

	return emailID, nil
}

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

// Create new Goblue client with default values
func NewClient(apiKey string) *Goblue {
	return &Goblue{
		ApiKey:            apiKey,
		Timeout:           DEFAULT_SEND_TIMEOUT,
		BaseUrl:           BASE_URL,
		ContentType:       CONTENT_TYPE,
		Method:            POST,
		ContentTypeHeader: CONTENT_TYPE_HEADER,
		ApiKeyHeader:      API_KEY_HEADER,
		EmailUrl:          EMAIL_URL,
		EmailTemplateUrl:  EMAIL_TEMPLATE_URL,
	}
}

// Send email
func (g *Goblue) SendEmail(email *Email) (*Response, error) {
	body := &bytes.Buffer{}
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(email)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.EmailUrl

	res, err := sendEmail(g.Method, urlStr, body)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(res.Body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Send email using template
func (g *Goblue) SendEmailTemplate(emailTemplate *EmailTemplate) (*Response, error) {
	body := &bytes.Buffer
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(emailTemplate)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.EmailTemplateUrl + "/" + strconv.Itoa(emailTemplate.Id)

	res, err := sendEmail(g.Method, urlStr, body)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(res.Body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendEmail(method string, url string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range email.Headers {
		req.Header.Add(key, val)
	}
	req.Header.Add(CONTENT_TYPE_HEADER, CONTENT_TYPE)
	req.Header.Add(API_KEY_HEADER, g.ApiKey)

	client := &http.Client{
		Timeout: g.Timeout,
	}

	return client.Do(req)
}
