# golang-echo-apigithub

This project is an API made in `Go` using the [Echo Web Framework](https://echo.labstack.com/) for building services, consuming the [GitHub API](https://docs.github.com/en/rest), and store data with [GormDB Library](https://gorm.io/) in the PostgreSQL database.

## Setup

```bash
$ docker-compose -f ./build/builder/docker-compose.yml up -d
```

## Run

```bash
$ docker-compose -f ./build/builder/docker-compose.yml exec service go run ./cmd/golang-echo-apigithub/main.go
```

## Test

Open in your browser or execute:
```bash
$ curl -i -X GET http://localhost:1323/rankings
```
