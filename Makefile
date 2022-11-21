run:
	go build -o bin/main main.go
	./bin/main serve

dep:
	go mod download
