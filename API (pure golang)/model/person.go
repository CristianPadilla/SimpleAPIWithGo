package model

//Person .
type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

// Persons slice of person
type Persons []Person
