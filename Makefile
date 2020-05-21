CC = go build
.PHONY: default build run
default: build run
build: main.go
	$(CC) main.go
run: main
	./main
main.go:
	$(error "main.go undefined")
