package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CristianPadilla/APIs/clase7-ECHO/authorizarion"
	"github.com/labstack/echo"
)

// Log middleware para que cada que me hagan una peticion a cualquiera de mis handler, se muestre la url y el metodo
// func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("peticion: %q, metodo: %q ", r.URL.Path, r.Method)
// 		f(w, r)
// 	}
// }

//TimeElapsed para ver el tiempo de ejecucion
func TimeElapsed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		defer fmt.Println("tiempo transcurrido: ", time.Since(inicio))
		log.Printf("peticion: %q, metodo: %q ", r.URL.Path, r.Method)
		f(w, r)
	}
}

//CheckAuthentication .
func CheckAuthentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Authorization es uno de los elementos del header
		token := c.Request().Header.Get("Authorization")
		_, err := authorizarion.ValidateToken(token)
		if err != nil {
			//responder que no esta permitido
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido el acceso"})
		}
		return f(c)
	}
}
