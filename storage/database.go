package storage

import (
	"fmt"
	"log"

	"github.com/RaeedAsif/flare-go-test/errors"
	"github.com/RaeedAsif/flare-go-test/models"
)

var (
	users []models.User
	SKIP  = 0
)

func Init() {
	users = make([]models.User, 0)

	//popualting dataset to users
	SKIP := 0
	total := LoadDataset(SKIP)

	log.Println(fmt.Sprintf("added %d users to storage", total))
}

func SetUser(user models.User) error {
	users = append(users, user)
	return nil
}

func GetUser(username string) (*models.User, error) {
	for _, u := range users {
		if username == u.Username {
			return &u, nil
		}
	}

	return nil, fmt.Errorf(errors.NOUSER)
}
