package goinblue

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const (
	TIME_FORMAT = "2013-06-20 20:09:22"
)

// The response from server
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SMSResponseData struct {
	Status          string            `json:"status"`
	Message         string            `json:"message"`
	NumberSent      int               `json:"number_sent"`
	To              string            `json:"to"`
	SMSCount        int               `json:"sms_count"`
	CreditsUsed     float64           `json:"credits_used"`
	RemainingCredit float64           `json:"remaining_credit"`
	Reference       map[string]string `json:"reference"`
	Description     string            `json:"description"`
	Reply           string            `json:"reply"`
	BounceType      string            `json:"bounce_type"`
	ErrorCode       int               `json:"error_code"`
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

func (r *Response) GetSMSResponseData() (*SMSResponseData, error) {
	smsResponse, ok := r.Data.(*SMSResponseData)
	if !ok {
		return nil, fmt.Errorf("Invalid Data type: ", "Not a valid SMSResponseData")
	}

	return smsResponse, nil
}

type WebhookResponse struct {
	Event         string    `json:"event"`
	Email         string    `json:"email"`
	Id            int64     `json:"id"`
	Date          time.Time `json:"date_time"`
	Ts            int64     `json:"ts"` // timestamp (same as Date but GMT)
	Subject       string    `json:"subject"`
	SendingIP     net.IP    `json:"sending-ip"`
	MessageId     string    `json:"message-id"`
	Tag           string    `json:"tag"`
	XMailinCustom string    `json:"X-Mailin-custom"`
	Reason        string    `json:"reason,omitempty"` // for "bounce" and "deferred" events only
	Link          string    `json:"link,omitempty"`   // for "click" events only
}

func (e *WebhookResponse) UnmarshalJSON(b []byte) error {
	type t struct {
		Event         string
		Email         string
		Id            int64
		Date          string
		Ts            int64
		Subject       string
		SendingIP     string `json:"sending_ip"`
		MessageId     string `json:"message-id"`
		Tag           string
		XMailinCustom string `json:"X-Mailin-custom"`
		Reason        string
		Link          string
	}
	var v t
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	date, err := time.Parse("2006-01-02 15:04:05", v.Date)
	if err != nil {
		return err
	}
	*e = WebhookResponse{
		Event:         v.Event,
		Email:         v.Email,
		Id:            v.Id,
		Date:          date,
		Ts:            v.Ts,
		Subject:       v.Subject,
		SendingIP:     net.ParseIP(v.SendingIP),
		MessageId:     v.MessageId,
		Tag:           v.Tag,
		XMailinCustom: v.XMailinCustom,
		Reason:        v.Reason,
		Link:          v.Link,
	}
	return nil
}
