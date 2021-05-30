package repositories

import (
	"go-api/config"
	"go-api/models"
	"go-api/schemas"
)

func GetAllBooks() (*[]models.Book, error) {
	var books *[]models.Book
	db := config.GetDatabase()
	err := db.Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

func GetBook(bookId string) (models.Book, error) {
	var book models.Book
	db := config.GetDatabase()
	err := db.First(&book, bookId).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func CreateBook(book_schema *schemas.CreateBookSchema) (models.Book, error) {
	db := config.GetDatabase()

	book := models.Book{
		Title:     book_schema.Title,
		Author:    book_schema.Author,
		Price:     book_schema.Price,
		Published: book_schema.Published,
	}
	err := db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func UpdateBook(bookId string, book_schema *schemas.CreateBookSchema) (models.Book, error) {
	var book models.Book
	db := config.GetDatabase()

	if errNotFound := db.First(&book, bookId).Error; errNotFound != nil {
		return book, errNotFound
	}

	book.Title = book_schema.Title
	book.Author = book_schema.Author
	book.Price = book_schema.Price
	book.Published = book_schema.Published

	if errNotUpdate := db.Updates(&book).Error; errNotUpdate != nil {
		return book, errNotUpdate
	}

	return book, nil
}

func DeleteBook(bookId string) error {
	db := config.GetDatabase()

	err := db.Delete(&models.Book{}, bookId).Error

	if err != nil {
		return err
	}

	return nil
}
