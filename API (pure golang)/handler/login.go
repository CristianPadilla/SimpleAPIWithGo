package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CristianPadilla/APIs/clase3-CRUD/authorizarion"
	"github.com/CristianPadilla/APIs/clase3-CRUD/model"
)

type login struct {
	storage Storage
}

func newLogin(storage Storage) login {
	return login{storage}
}

func (p *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed Method", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	//save login data at structure
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "login does not have correct structure", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	if !loginIsValid(data) {
		response := newResponse(Error, "incorrect user or pasword", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	token, err := authorizarion.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "error generating token", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	tokenData := map[string]string{"token": token}
	response := newResponse(Message, "OK", tokenData)
	ResponseJSON(w, http.StatusOK, response)
}

//simulation of login
func loginIsValid(data model.Login) bool {
	return data.Email == "@mail.com" && data.Password == "12345"
}
