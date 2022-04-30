# Langi backend


## Description

The langi app can you help to create dictionaries and words to learn foreign languages. 

* REST API 
* JWT-based authorization and authentication
* Clean architecture
* Postgresql
* Dockerized
* Swagger documentation
* Goland Http client for tests api


## Init and Run

1. Add password in DB_PASSWORD variable to `.env` file and duplicate it in `docker-compose.yml` in `langi` and `db` sections 
2. `make init && make up` to download deps, init and run docker containers and starting the app  

* use `make up` in second running's, additional commands you'll see in Makefile
* `http` folder contains examples of the requests to API. First, create a user call Registration method, next call Login. Token is will be saved into a global variable. Later you can run any another methods in API, create dictionary or words for example.
* `make test` to start tests 