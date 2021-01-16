package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Person .
type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

func CreatePerson(url, token string, p *Person) GeneralResponse {

	data := bytes.NewBuffer([]byte{})
	//take person struct, become it to json and put into data
	err := json.NewEncoder(data).Encode(&p)
	if err != nil {
		log.Fatalf("Error in marshal of persona %v", err)
	}
	//execute petition
	response := httpClient(http.MethodPost, url, token, data)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}
	if response.StatusCode != http.StatusCreated {
		log.Fatalf("expected code 200, got: %d ", response.StatusCode)

	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("error at unmarshal of body at login: %d", err)
	}
	return dataResponse

}
