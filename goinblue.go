// Golang library for sendinblue API (work in progress)

package goinblue

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	SMS_URL             = "/sms"
	POST                = "POST"
	CONTENT_TYPE_HEADER = "Content-Type"
	API_KEY_HEADER      = "api-key"
)

var (
	DEFAULT_SEND_TIMEOUT time.Duration = time.Second * 30
)

// The main struct of this package
type Goinblue struct {
	ApiKey            string
	Timeout           time.Duration
	BaseUrl           string
	ContentType       string
	Method            string
	ContentTypeHeader string
	ApiKeyHeader      string
	EmailUrl          string
	EmailTemplateUrl  string
	SMSUrl            string
}

// Create new Goinblue client with default values
func NewClient(apiKey string) *Goinblue {
	return &Goinblue{
		ApiKey:            apiKey,
		Timeout:           DEFAULT_SEND_TIMEOUT,
		BaseUrl:           BASE_URL,
		ContentType:       CONTENT_TYPE,
		Method:            POST,
		ContentTypeHeader: CONTENT_TYPE_HEADER,
		ApiKeyHeader:      API_KEY_HEADER,
		EmailUrl:          EMAIL_URL,
		EmailTemplateUrl:  EMAIL_TEMPLATE_URL,
		SMSUrl:            SMS_URL,
	}
}

// Send email
func (g *Goinblue) SendEmail(email *Email) (*Response, error) {
	body := &bytes.Buffer{}
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(email)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.EmailUrl

	res, err := g.sendMessage(g.Method, urlStr, email.Headers, ioutil.NopCloser(body), body.Len())
	if err != nil {
		return nil, err
	}
	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Failed to send email: %s", res.Status)
	}

	resp := &Response{}
	err = json.Unmarshal(rawResBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Send email using template
func (g *Goinblue) SendEmailTemplate(emailTemplate *EmailTemplate) (*Response, error) {
	body := &bytes.Buffer{}
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(emailTemplate)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.EmailTemplateUrl + "/" + strconv.Itoa(emailTemplate.Id)

	res, err := g.sendMessage(g.Method, urlStr, emailTemplate.Headers, ioutil.NopCloser(body), body.Len())
	if err != nil {
		return nil, err
	}
	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Failed to send email: %s", res.Status)
	}

	resp := &Response{}
	err = json.Unmarshal(rawResBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *Goinblue) SendSMS(sms *SMS) (*Response, error) {
	body := &bytes.Buffer{}
	defer body.Reset()

	encoder := json.NewEncoder(body)
	err := encoder.Encode(sms)
	if err != nil {
		return nil, err
	}

	urlStr := g.BaseUrl + g.SMSUrl

	res, err := g.sendMessage(g.Method, urlStr, nil, ioutil.NopCloser(body), body.Len())
	if err != nil {
		return nil, err
	}
	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	rawResBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Failed to send sms: %s", res.Status)
	}

	resp := &Response{}
	err = json.Unmarshal(rawResBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *Goinblue) sendMessage(method string, url string, headers map[string]string, body io.ReadCloser, length int) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(length)

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
