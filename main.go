package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RaeedAsif/flare-go-test/routes"
	"github.com/RaeedAsif/flare-go-test/storage"
)

func main() {
	PORT := "8080"

	//Init routers
	router := routes.Init()
	log.Println("routers intiliased")

	//Init Database
	storage.Init()
	log.Println("storage intiliased")

	//create http server
	fmt.Println("listening http server start on " + string(PORT))
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
