# golang-echo-apigithub

A estrutura desse projeto foi feita baseada em [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

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
$ curl -i -X GET http://localhost:1323/rankings/100
```