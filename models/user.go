package models

// User struct
type User struct {
	ID        int    `json:"id"`        // unique id
	FirstName string `json:"firstName"` // first name of user
	LastName  string `json:"lastName"`  // last name of user
	Email     string `json:"email"`     // email of user
	Username  string `json:"username"`  // username of user
}
