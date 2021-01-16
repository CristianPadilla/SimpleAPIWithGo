package main

import (
	"log"

	"github.com/CristianPadilla/APIs/clase7-ECHO/authorizarion"
	"github.com/CristianPadilla/APIs/clase7-ECHO/handler"
	"github.com/CristianPadilla/APIs/clase7-ECHO/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := authorizarion.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("can not load certificates: %v", err)
	}
	store := storage.NewMemory()
	e := echo.New()
	// middleware to prevent server crash
	e.Use(middleware.Recover())
	// middleware to register every petition
	e.Use(middleware.Logger())
	// "store" can be replaced for any SQL storage system like MySQL
	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)
	log.Println("servidor iniciado")
	//execute server
	err = e.Start(":8080")
	if err != nil {
		log.Println("error en el servidor")
	}
}
