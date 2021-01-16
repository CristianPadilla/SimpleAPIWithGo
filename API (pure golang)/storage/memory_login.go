package storage

//MemLogin .
type MemLogin struct {
	Email    string
	Password string
}

//NewMemLogin constructor method
func NewMemLogin() MemLogin {
	return MemLogin{}
}
