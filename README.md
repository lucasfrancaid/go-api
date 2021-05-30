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
[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=Go%20API&uri=https%3A%2F%2Fgithub.com%2Flucasfrancaid%2Fgo-api%2Fblob%2Fmain%2Fdocs%2Finsomnia.json)

### Authors
[POST] Create Author --> http://localhost:8080/v1/authors  
[GET] Get All Authors --> http://localhost:8080/v1/authors  
[GET] Get Author --> http://localhost:8080/v1/authors/1  
[PUT] Update Author --> http://localhost:8080/v1/books/1  
[DELETE] Delete Authors --> http://localhost:8080/v1/authors/1  
```json
// [post, put] request schema
{
	"name": "Robert C. Martin (Uncle Bob)",
	"site": "https://github.com/unclebob"
}
```

### Books
[POST] Create Book --> http://localhost:8080/v1/books  
[GET] Get All Books --> http://localhost:8080/v1/books  
[GET] Get Book --> http://localhost:8080/v1/books/1  
[PUT] Update Book --> http://localhost:8080/v1/books/1  
[DELETE] Delete Books --> http://localhost:8080/v1/books/1  
```json
// [post, put] request schema
{
	"title": "Clean Coder",
	"author_id": 1,
	"price": 19.89,
	"published": "2018-11-22T15:04:00Z"
}
```
