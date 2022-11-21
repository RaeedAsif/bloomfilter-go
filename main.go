package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"

	"github.com/RaeedAsif/flare-go-test/routes"
	"github.com/RaeedAsif/flare-go-test/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}

	PORT := os.Getenv("PORT")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	//Init Database
	storage.Init()
	log.Println("storage intiliased")

	//Init routers
	router := routes.Init()
	log.Println("routers intiliased")

	srv := &http.Server{
		Addr: "0.0.0.0:" + PORT,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	//create http server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("listening http server start on :" + string(PORT))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
