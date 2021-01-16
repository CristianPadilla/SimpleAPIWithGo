package model

import "errors"

var (
	//ErrPersonCanNotBeNull .
	ErrPersonCanNotBeNull error = errors.New("La persona no puede ser nula")
	ErrPersonDoesNotExist error = errors.New("La que intenta buscar no existe")
)
