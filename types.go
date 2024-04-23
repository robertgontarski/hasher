package main

import "github.com/robertgontarski/hasher/proto"

type Topic string

type ExtraDataHasher map[string]any

type EmailChangeMessage struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type PhoneChangeMessage struct {
	ID          int    `json:"id"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

type NameChangeMessage struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type ResponseMap map[string]any

type ResponseEndpoint struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    ResponseMap `json:"data"`
}

func NewResponseEndpoint(status int, msg string, data ResponseMap) *ResponseEndpoint {
	return &ResponseEndpoint{
		Status:  status,
		Message: msg,
		Data:    data,
	}
}

type HasherResponseCh struct {
	Data *proto.HashResponse
	Err  error
}
