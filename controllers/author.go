package controllers

import (
	"go-api/repositories"
	"go-api/schemas"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(ctx *gin.Context) {
	authors, err := repositories.GetAllAuthors()

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, authors)
}

func GetAuthor(ctx *gin.Context) {
	authorId := ctx.Param("id")
	author, err := repositories.GetAuthor(authorId)

	if err != nil {
		ctx.JSON(404, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, author)
}

func CreateAuthor(ctx *gin.Context) {
	var author_schema *schemas.CreateAuthorSchema

	if err := ctx.ShouldBindJSON(&author_schema); err != nil {
		ctx.JSON(400, gin.H{"detail": err})
		return
	}

	author, err := repositories.CreateAuthor(author_schema)

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, author)
}

func UpdateAuthor(ctx *gin.Context) {
	var author_schema *schemas.CreateAuthorSchema
	authorId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&author_schema); err != nil {
		ctx.JSON(400, gin.H{"detail": err})
		return
	}

	author, err := repositories.UpdateAuthor(authorId, author_schema)

	if err != nil {
		ctx.JSON(400, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, author)
}

func DeleteAuthor(ctx *gin.Context) {
	authorId := ctx.Param("id")
	err := repositories.DeleteAuthor(authorId)

	if err != nil {
		ctx.JSON(404, gin.H{"detail": err.Error()})
		return
	}

	ctx.JSON(200, "Ok")
}
