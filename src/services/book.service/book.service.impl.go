package bookservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/src/data/response"
	bookrepository "github.com/kunnoh/lms-api/src/repository/book.repository"
)

type BookServiceImpl struct {
	BookRepo bookrepository.BookRepository
	validate *validator.Validate
}

func NewBookServiceImpl(bookRepository bookrepository.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepo: bookRepository,
		validate: validate,
	}
}

// Create implements BookService.
func (b *BookServiceImpl) Create() response.Response {
	panic("unimplemented")
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(BookId string) response.Response {
	panic("unimplemented")
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll() response.Response {
	panic("unimplemented")
}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(BookId string) response.Response {
	panic("unimplemented")
}

// Update implements BookService.
func (b *BookServiceImpl) Update() response.Response {
	panic("unimplemented")
}
