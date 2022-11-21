package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/RaeedAsif/flare-go-test/models"
)

// dataset:  https://dummyjson.com/users

// Dataset struct
type Dataset struct {
	Users []models.User `json:"users"` // list of users
	Limit int           `json:"limit"` // api limit
	Skip  int           `json:"skip"`  // api skip
	Total int           `json:"total"` // total items in db
}

var LIMIT = 0

// LoadDataset fetches and loads the api data to storage db
func LoadDataset(skip int) int {
	url := "https://dummyjson.com/users"

	if skip > 0 {
		url = "https://dummyjson.com/users?skip=" + fmt.Sprint(skip)
	}

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	dataset := Dataset{}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &dataset)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	users = append(users, dataset.Users...)

	LIMIT = LIMIT + dataset.Limit

	if dataset.Total > len(users) {
		return LoadDataset(LIMIT)
	}

	return LIMIT
}
