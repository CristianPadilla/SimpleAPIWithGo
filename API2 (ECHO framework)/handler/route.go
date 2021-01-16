package handler

import (
	"github.com/CristianPadilla/APIs/clase7-ECHO/middleware"
	"github.com/labstack/echo"
)

//RoutePerson .
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	// person handler group
	persons := e.Group("/v1/persons")
	persons.Use(middleware.CheckAuthentication)
	persons.POST("", h.create)
	persons.PUT("/:id", h.update)
	persons.DELETE("/:id", h.delete)
	persons.GET("/:id", h.getByID)
	persons.GET("", h.getAll)

}

//RouteLogin .
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
