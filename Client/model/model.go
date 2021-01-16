package model

import (
	"io"
	"log"
	"net/http"
)

type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

//method for making petitions
func httpClient(method, url, token string, body io.Reader) *http.Response {
	//create petition
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	//execute petition
	response, err := client.Do(req)
	if err != nil {
		log.Printf("request: %v", err)
	}

	return response
}
