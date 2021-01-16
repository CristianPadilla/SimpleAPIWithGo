package storage

import (
	"fmt"

	"github.com/CristianPadilla/APIs/clase7-ECHO/model"
)

//Memory .
type Memory struct {
	currentID uint
	Persons   map[uint]model.Person
}

//NewMemory constructor method
func NewMemory() Memory {
	persons := make(map[uint]model.Person)
	return Memory{
		currentID: 0,
		Persons:   persons,
	}

}

//Create .
func (m *Memory) Create(p *model.Person) error {
	if p == nil {
		return model.ErrPersonCanNotBeNull
	}
	m.currentID++
	m.Persons[m.currentID] = *p

	return nil

}

//Update .
func (m *Memory) Update(ID uint, p *model.Person) error {
	if p == nil {
		return model.ErrPersonCanNotBeNull
	}
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID :%d %w", ID, model.ErrPersonDoesNotExist)
	}
	m.Persons[ID] = *p
	return nil
}

//Delete .
func (m *Memory) Delete(ID uint) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID :%d %w", ID, model.ErrPersonDoesNotExist)
	}
	delete(m.Persons, ID)
	return nil
}

//GetByID .
func (m *Memory) GetByID(ID uint) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID :%d %w", ID, model.ErrPersonDoesNotExist)
	}
	return person, nil
}

//GetAll .
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}
	return result, nil
}
