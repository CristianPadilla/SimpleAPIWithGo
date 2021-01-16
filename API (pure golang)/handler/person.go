package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/CristianPadilla/APIs/clase3-CRUD/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	//verify method
	if r.Method != http.MethodPost {
		response := newResponse(Error, "not allowed Method", nil)
		//execute json response
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "person does not have correct structure", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "there was a problem creating person", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	//finaly
	response := newResponse(Message, "OK", nil)
	ResponseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "not allowed Method", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "person does not have correct structure", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Update(uint(ID), &data)
	if err != nil {
		response := newResponse(Error, "error at modify person", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Error, "OK", nil)
	ResponseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "not allowed Method", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Delete(uint(ID))
	if errors.Is(err, model.ErrPersonDoesNotExist) {
		response := newResponse(Error, "ID does nor exist", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "error deleting record", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Error, "OK", nil)
	ResponseJSON(w, http.StatusOK, response)
}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "not allowed Method", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	data, err := p.storage.GetByID(uint(ID))
	if errors.Is(err, model.ErrPersonDoesNotExist) {
		response := newResponse(Error, "ID does nor exist", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "error getting record", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Error, "OK", data)
	ResponseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "not allowed Method", nil)
		ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "error getting records ", nil)
		ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	ResponseJSON(w, http.StatusOK, response)

}
