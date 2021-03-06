package handler

import "github.com/CristianPadilla/APIs/clase3-CRUD/model"

type Storage interface {
	Create(p *model.Person) error
	Update(id uint, p *model.Person) error
	Delete(ID uint) error
	GetByID(ID uint) (model.Person, error)
	GetAll() (m model.Persons, e error)
}
