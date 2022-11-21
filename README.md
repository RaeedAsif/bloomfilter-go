# flare-go-test
This is a implementation of microservice checking if user exists using a data structure Bloom Filter and Go Mux

## Preface
This repository is the code repo of microservice to check the availability of a username using golang.
This repository implements and uses Bloom filter.
This sample uses [Go mux](https://github.com/gorilla/mux) as web application framework.
This sample application provides only following functions as Web APIs.
  1. "/health" : checks for server health
  2. "/user/username/{username}" : checks if username exists
  3. "/docs" : swagger docs

## Install
Perform the following steps:
1. Download and install [Visual Studio Code(VS Code)](https://code.visualstudio.com/).
2. Download and install [Golang](https://golang.org/).

## Run
Perform the following steps:
1. Rename the file "env_sample" to "env". 
2. Run "make run" in directory terminal to run the microservice at 8080
3. Run "make dep" in directory terminal to download and install required go dependencies
3. open "localhost:8080/docs" on browser.
