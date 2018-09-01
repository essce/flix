all: 
build:
	go build -o bin/flix.bin cmd/main.go
run: 
	go build -o bin/flix.bin cmd/main.go
	bin/flix.bin