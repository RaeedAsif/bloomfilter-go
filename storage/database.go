package storage

import (
	"fmt"
	"log"

	"github.com/RaeedAsif/flare-go-test/models"
)

var (
	users []models.User
	SKIP  = 0
)

// Init initiates storage db
func Init() {
	users = make([]models.User, 0)

	log.Println("fetching users from https://dummyjson.com/users")

	//popualting dataset to users
	SKIP := 0
	total := LoadDataset(SKIP)

	// bloom filter init
	InitBloomFilter()
	bf := GetInstance()

	// set bitset
	for _, u := range users {
		bf.Set(u.Username)
	}

	log.Println(fmt.Sprintf("fetched and added %d users to storage", total))
}

// SetUser sets in memory user db
func SetUser(user models.User) error {
	bf := GetInstance()
	bf.mu.Lock()

	users = append(users, user)

	bf.mu.Unlock()

	return nil
}

// IsUsernameExists checks if user exists in db
func IsUsernameExists(username string) bool {

	bf := GetInstance()

	if bf.Check(username) {
		return true
	}

	for _, u := range users {
		if username == u.Username {
			return true
		}
	}

	return false
}
