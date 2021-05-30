# Go Api
My first API developed in Go, built with Gin and Gorm.

## Requirements
- Go 1.16
- Docker & Docker Compose

## Setup
Run below code to setup go-api with docker:
```zsh
docker-compose --build up
# or
make run_docker
```

If you want use go-api without docker:
```zsh
make run_postgres # you need a postgres database
make run
```

## API Requests
[GET] Get All Books - http://localhost:8080/books

[GET] Get Book - http://localhost:8080/books/1

[POST] Create Book - http://localhost:8080/books
```json
{
	"title": "Clean Coder",
	"author": "Robert C. Martin (Uncle Bob)",
	"price": 19.89,
	"published": "2018-11-22T15:04:00Z"
}
```

[PUT] Update Book - http://localhost:8080/books/1
```json
{
	"title": "Clean Architecture",
	"author": "Robert C. Martin (Uncle Bob)",
	"price": 34.89,
	"published": "2019-12-05T12:04:00Z"
}
```

[DELETE] Delete Books - http://localhost:8080/books/1
