package main

import (
	"log"
	"net/http"

	"github.com/CristianPadilla/APIs/clase3-CRUD/authorizarion"
	"github.com/CristianPadilla/APIs/clase3-CRUD/handler"
	"github.com/CristianPadilla/APIs/clase3-CRUD/storage"
)

func main() {
	//load certificates
	err := authorizarion.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("couldn't load certificates: %v", err)
	}

	// initialize storage system
	store := storage.NewMemory()
	//initialize router
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)
	log.Println("server started")
	//ejecutar servidor
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("error at server")
	}
}
