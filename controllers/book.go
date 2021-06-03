package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasfrancaid/go-api/repositories"
	"github.com/lucasfrancaid/go-api/schemas"
)

func GetAllBooks(ctx *gin.Context) {
	books, err := repositories.GetAllBooks()

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, books)
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	book, err := repositories.GetBook(bookId)

	if err != nil {
		ctx.JSON(404, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, book)
}

func CreateBook(ctx *gin.Context) {
	var book_schema *schemas.CreateBookSchema

	if err := ctx.ShouldBindJSON(&book_schema); err != nil {
		ctx.JSON(400, gin.H{"detail": err})
		return
	}

	book, err := repositories.CreateBook(book_schema)

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, book)
}

func UpdateBook(ctx *gin.Context) {
	var book_schema *schemas.CreateBookSchema
	bookId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&book_schema); err != nil {
		ctx.JSON(400, gin.H{"detail": err})
		return
	}

	book, err := repositories.UpdateBook(bookId, book_schema)

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, book)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	err := repositories.DeleteBook(bookId)

	if err != nil {
		ctx.JSON(404, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, "Ok")
}
