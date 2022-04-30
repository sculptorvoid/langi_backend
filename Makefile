deps:
	go mod download

build:
	docker-compose build langi

up:
	docker-compose up langi

init: deps build

test:
	go test -v ./...

swag_init:
	swag init -g cmd/main.go

swag_fmt:
	swag fmt -g cmd/main.go