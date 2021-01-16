package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CristianPadilla/APIs/clase3-CRUD/authorizarion"
)

// Log middleware to register every server petition
func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion: %q, metodo: %q ", r.URL.Path, r.Method)
		f(w, r)
	}
}

//TimeElapsed middleware to register how long takes every petition
func TimeElapsed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		defer fmt.Println("tiempo transcurrido: ", time.Since(inicio))
		log.Printf("peticion: %q, metodo: %q ", r.URL.Path, r.Method)
		f(w, r)
	}
}

//CheckAuthentication .
func CheckAuthentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authorization es uno de los elementos del header
		token := r.Header.Get("Authorization")
		_, err := authorizarion.ValidateToken(token)
		if err != nil {
			//responder que no esta permitido
			forbbiden(w, r)
			return
		}
		f(w, r)
	}
}

//helper method
func forbbiden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("not authorized"))
}
