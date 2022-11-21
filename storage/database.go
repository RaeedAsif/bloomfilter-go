package storage

import (
	"fmt"
	"log"

	"github.com/RaeedAsif/flare-go-test/errors"
	"github.com/RaeedAsif/flare-go-test/models"
	"github.com/RaeedAsif/flare-go-test/utils/bloom"
)

var (
	users []models.User
	SKIP  = 0
)

func Init() {
	users = make([]models.User, 0)

	log.Println("fetching users from https://dummyjson.com/users")

	//popualting dataset to users
	SKIP := 0
	total := LoadDataset(SKIP)

	// bloom filter init
	bloom.Init()
	bf := bloom.GetInstance()

	for _, u := range users {
		bf.Set(u.Username)
	}

	log.Println(fmt.Sprintf("fetched and added %d users to storage", total))
}

func SetUser(user models.User) error {
	users = append(users, user)
	return nil
}

func GetUser(username string) (*models.User, error) {

	bf := bloom.GetInstance()

	if !bf.Check(username) {
		fmt.Println("HERE")
		return nil, fmt.Errorf(errors.ERR_NO_USER)
	}

	for _, u := range users {
		if username == u.Username {
			return &u, nil
		}
	}

	return nil, fmt.Errorf(errors.ERR_NO_USER)
}
