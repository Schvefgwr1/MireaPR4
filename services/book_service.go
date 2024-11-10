package services

import (
	"MireaPR4/models"
	"MireaPR4/repositories"
)

type BookService interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBooks() ([]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id string) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(book *models.Book) (*models.Book, error) {
	if err := s.repo.Create(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id string) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
