package handler

import (
	"net/http"

	"github.com/CristianPadilla/APIs/clase7-ECHO/authorizarion"
	"github.com/CristianPadilla/APIs/clase7-ECHO/model"
	"github.com/labstack/echo"
)

type login struct {
	storage Storage
}

func newLogin(storage Storage) login {
	return login{storage}
}

func (p *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "login does not have correct structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if !loginIsValid(data) {
		response := newResponse(Error, "incorrect user or pasword", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	token, err := authorizarion.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "error generating token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	tokenData := map[string]string{"token": token}
	response := newResponse(Message, "OK", tokenData)
	return c.JSON(http.StatusOK, response)
}

func loginIsValid(data model.Login) bool {
	return data.Email == "@mail.com" && data.Password == "12345"
}
