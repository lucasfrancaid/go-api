package repositories

import (
	"github.com/lucasfrancaid/go-api/config"
	"github.com/lucasfrancaid/go-api/models"
	"github.com/lucasfrancaid/go-api/schemas"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	db := config.GetDatabase()

	if err := db.Find(&books).Error; err != nil {
		return books, err
	}

	if len(books) > 0 {
		for index, _ := range books {
			if err := db.Model(&models.Author{}).Where("id = ?", books[index].AuthorID).Take(&books[index].Author).Error; err != nil {
				return books, err
			}
		}
	}

	return books, nil
}

func GetBook(bookId string) (models.Book, error) {
	var book models.Book
	db := config.GetDatabase()

	if err := db.First(&book, bookId).Error; err != nil {
		return book, err
	}

	if book.ID != 0 {
		if err := db.Model(&models.Author{}).Where("id = ?", book.AuthorID).Take(&book.Author).Error; err != nil {
			return book, err
		}
	}

	return book, nil
}

func CreateBook(book_schema *schemas.CreateBookSchema) (models.Book, error) {
	db := config.GetDatabase()

	book := models.Book{
		Title:     book_schema.Title,
		AuthorID:  book_schema.AuthorID,
		Price:     book_schema.Price,
		Published: book_schema.Published,
	}

	if err := db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func UpdateBook(bookId string, book_schema *schemas.CreateBookSchema) (models.Book, error) {
	var book models.Book
	db := config.GetDatabase()

	if errBookNotFound := db.First(&book, bookId).Error; errBookNotFound != nil {
		return book, errBookNotFound
	}

	book.Title = book_schema.Title
	book.AuthorID = book_schema.AuthorID
	book.Price = book_schema.Price
	book.Published = book_schema.Published

	if errNotUpdate := db.Updates(&book).Error; errNotUpdate != nil {
		return book, errNotUpdate
	}

	return book, nil
}

func DeleteBook(bookId string) error {
	db := config.GetDatabase()

	if err := db.Delete(&models.Book{}, bookId).Error; err != nil {
		return err
	}

	return nil
}
