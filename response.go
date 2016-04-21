package goblue

import (
	"encoding/json"
	"fmt"
	"time"
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
	Id            int       `json:"id"`
	Date          time.Time `json:"date"`
	MessageId     string    `json:"message-id"`
	Tag           string    `json:"tag"`
	XMailinCustom string    `json:"X-Mailin-custom"`
	Reason        string    `json:"reason"`
	Link          string    `json:"link"`
}

func UnmarshalWebhookResponse(data []byte) (*WebhookResponse, error) {
	res := &WebhookResponse{}

	err := json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
