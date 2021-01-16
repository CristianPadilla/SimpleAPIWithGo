package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/CristianPadilla/APIs/clase7-ECHO/model"
	"github.com/labstack/echo"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person does not have correct structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "there was a problem creating person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	//en caso de que todo salga bien...
	response := newResponse(Message, "OK", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person does not have correct structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Update(uint(ID), &data)
	if err != nil {
		response := newResponse(Error, "error at modify person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Error, "OK", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Delete(uint(ID))
	if errors.Is(err, model.ErrPersonDoesNotExist) {
		response := newResponse(Error, "ID does nor exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "error deleting record", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID is not valid", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data, err := p.storage.GetByID(uint(ID))
	if errors.Is(err, model.ErrPersonDoesNotExist) {
		response := newResponse(Error, "ID does nor exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Error getting record", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Error, "OK", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "error getting records", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", data)
	return c.JSON(http.StatusOK, response)

}
