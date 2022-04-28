# Langi backend


## Description

Create your own unlimited dictionaries and learn languages or any other things

* REST API 
* JWT-based authorization and authentication
* Clean architecture
* Postgresql 
* Swagger documentation

## Init

1. create .env file in the root folder
2. add to .env key DB_PASSWORD="with_your_password_from_postgress"
3. `docker pull postgres`
4. `docker run --name=langi-app-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres`
5. `migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up`
6. `go run cmd/main.go`