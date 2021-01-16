package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error at login marshal  %v", err)
	}
	response := httpClient(http.MethodPost, url, "", data)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("reading error in body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("expected code 200, got : %d", response.StatusCode)
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error in unmarshal of person's body: %d", err)
	}
	return dataResponse

}
