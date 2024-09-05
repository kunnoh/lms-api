package bookrepository

import "github.com/kunnoh/lms-api/src/model"

type BookRepository interface {
	Save(book model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(bookId string) error
	Search(s string) (books []model.Book, err error)
	FindById(bookId string) (book model.Book, err error)
	FindAll() (books []model.Book, err error)
}
