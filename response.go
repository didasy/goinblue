package goblue

import (
	"fmt"
)

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
