package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	PORT := "8080"

	//routers init
	router := mux.NewRouter()

	//create http server
	fmt.Println("listening http server start on " + string(PORT))
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
