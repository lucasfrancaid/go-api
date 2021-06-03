package repositories

import (
	"github.com/lucasfrancaid/go-api/config"
	"github.com/lucasfrancaid/go-api/models"
	"github.com/lucasfrancaid/go-api/schemas"
)

func GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	db := config.GetDatabase()

	if err := db.Find(&authors).Error; err != nil {
		return authors, err
	}

	return authors, nil
}

func GetAuthor(authorId string) (models.Author, error) {
	var author models.Author
	db := config.GetDatabase()

	if err := db.First(&author, authorId).Error; err != nil {
		return author, err
	}

	return author, nil
}

func GetAuthorBooks(authorId string) ([]models.Book, error) {
	var books []models.Book
	db := config.GetDatabase()

	if errAuthorNotFound := db.First(&models.Author{}, authorId).Error; errAuthorNotFound != nil {
		return books, errAuthorNotFound
	}

	if err := db.Where("author_id = ?", authorId).Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func CreateAuthor(author_schema *schemas.CreateAuthorSchema) (models.Author, error) {
	db := config.GetDatabase()

	author := models.Author{
		Name: author_schema.Name,
		Site: author_schema.Site,
	}

	if err := db.Create(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

func UpdateAuthor(authorId string, author_schema *schemas.CreateAuthorSchema) (models.Author, error) {
	var author models.Author
	db := config.GetDatabase()

	if errNotFound := db.First(&author, authorId).Error; errNotFound != nil {
		return author, errNotFound
	}

	author.Name = author_schema.Name
	author.Site = author_schema.Site

	if errNotUpdate := db.Updates(&author).Error; errNotUpdate != nil {
		return author, errNotUpdate
	}

	return author, nil
}

func DeleteAuthor(authorId string) error {
	db := config.GetDatabase()

	if err := db.Delete(&models.Author{}, authorId).Error; err != nil {
		return err
	}

	return nil
}
