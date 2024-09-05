package bookrepository

import "github.com/kunnoh/lms-api/src/model"

type BookRepository interface {
	Save(book model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(bookId string) error
	Search(s string) ([]model.Book, error)
	FindById(bookId string) (model.Book, error)
	FindAll() ([]model.Book, error)
}
