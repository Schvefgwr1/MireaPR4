package repositories

import (
	"MireaPR4/models"
	"errors"
	"log"
	"strconv"
	//"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id string) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id string) error
}

type bookRepository struct {
	books  []models.Book
	nextID int
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		books: []models.Book{
			{ID: "1", Title: "Go Programming", Author: "John Doe"},
			{ID: "2", Title: "Learning PostgreSQL", Author: "Jane Smith"},
			{ID: "3", Title: "Microservices in Go", Author: "Alex Johnson"},
		},
		nextID: 4,
	}
}

func (b *bookRepository) Create(book *models.Book) error {
	book.ID = strconv.Itoa(b.nextID)
	b.nextID++

	log.Println(book)
	b.books = append(b.books, *book)
	log.Println(b.books)
	return nil
}

func (b *bookRepository) GetAll() ([]models.Book, error) {
	log.Println("Books from getAll: ", b.books)
	return b.books, nil
}

func (b *bookRepository) GetByID(id string) (*models.Book, error) {
	for _, book := range b.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (b *bookRepository) Update(book *models.Book) error {
	for i, existingBook := range b.books {
		if existingBook.ID == book.ID {
			b.books[i] = *book
			return nil
		}
	}
	return errors.New("book not found")
}

func (b *bookRepository) Delete(id string) error {
	for i, existingBook := range b.books {
		if existingBook.ID == id {
			b.books = append(b.books[:i], b.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

//type bookRepository struct {
//	db *gorm.DB
//}
//
//func NewBookRepository(db *gorm.DB) BookRepository {
//	return &bookRepository{db}
//}
//
//func (r *bookRepository) Create(book *models.Book) error {
//	return r.db.Create(book).Error
//}
//
//func (r *bookRepository) GetAll() ([]models.Book, error) {
//	var books []models.Book
//	err := r.db.Find(&books).Error
//	return books, err
//}
//
//func (r *bookRepository) GetByID(id string) (*models.Book, error) {
//	var book models.Book
//	err := r.db.First(&book, "id = ?", id).Error
//	return &book, err
//}
//
//func (r *bookRepository) Update(book *models.Book) error {
//	return r.db.Save(book).Error
//}
//
//func (r *bookRepository) Delete(id string) error {
//	return r.db.Delete(&models.Book{}, "id = ?", id).Error
//}
