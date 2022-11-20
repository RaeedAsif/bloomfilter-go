package storage

import "fmt"

type User struct {
	Name     string
	Email    string
	Username string
}

var users []User

func InitUsers() {
	users = make([]User, 0)
}

func SetUser(user User) error {
	users = append(users, user)
	return nil
}

func GetUser(username string) (*User, error) {
	for _, u := range users {
		if username == u.Username {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("no user found")
}
