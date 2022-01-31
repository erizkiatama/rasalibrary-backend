package models

import (
	"fmt"
	"log"
)

type CustomError interface {
	LogError()
	GetStatus() int
}

type ClientError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func (ths *ClientError) Error() string {
	return fmt.Sprintf("%s: %s", ths.Code, ths.Message)
}

func (ths *ClientError) GetStatus() int {
	return ths.Status
}

func (ths *ClientError) LogError() {
	// do nothing
}

func NewClientError(code, message string, statusCode int) error {
	return &ClientError{
		Code:    code,
		Message: message,
		Type:    "client",
		Status:  statusCode,
	}
}

type ServerError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	originalErr error
}

func (ths *ServerError) Error() string {
	return fmt.Sprintf("%s - %s", ths.Code, ths.originalErr.Error())
}

func (ths *ServerError) LogError() {
	log.Println("ERROR: " + ths.Error())
}

func (ths *ServerError) GetStatus() int {
	return ths.Status
}

func NewServerError(code string, statusCode int, origin error) error {
	return &ServerError{
		Code:        code,
		Message:     "Internal Server Error",
		Type:        "server",
		Status:      statusCode,
		originalErr: origin,
	}
}
