package main

import (
	"fmt"

	"github.com/CristianPadilla/APIs/clase8-CLIENTE/model"
)

const url = "http://localhost:8080"

func main() {
	lc := model.LoginClient(url+"/v1/login", "@mail.com", "12345")
	// fmt.Print(lc)
	p := model.Person{
		Name: "cris",
		Age:  20,
		Communities: []model.Community{
			{Name: "software"},
			{Name: "golang"},
		},
	}
	gr := model.CreatePerson(url+"/v1/persons", lc.Data.Token, &p)
	fmt.Println(gr)
}
