package models

// User
type User struct {
	// id
	ID int `json:"id"`
	// firstName
	FirstName string `json:"firstName"`
	// lastName
	LastName string `json:"lastName"`
	// email
	Email string `json:"email"`
	// username
	Username string `json:"username"`
}
