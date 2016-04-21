// Golang library for sendinblue API

package goblue

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

	res, err := g.sendEmail(g.Method, urlStr, email.Headers, ioutil.NopCloser(body))
	if err != nil {
		return nil, err
	}

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rawResBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Send email using template
func (g *Goblue) SendEmailTemplate(emailTemplate *EmailTemplate) (*Response, error) {
	body := &bytes.Buffer{}
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(emailTemplate)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.EmailTemplateUrl + "/" + strconv.Itoa(emailTemplate.Id)

	res, err := g.sendEmail(g.Method, urlStr, emailTemplate.Headers, ioutil.NopCloser(body))
	if err != nil {
		return nil, err
	}

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rawResBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *Goblue) sendEmail(method string, url string, headers map[string]string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}
	req.Header.Add(g.ContentTypeHeader, g.ContentType)
	req.Header.Add(g.ApiKeyHeader, g.ApiKey)

	client := &http.Client{
		Timeout: g.Timeout,
	}

	return client.Do(req)
}
