package bookrepository

import (
	"github.com/kunnoh/lms-api/src/model"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return &BookServiceImpl{Db: Db}
}

// Delete implements BookRepository.
func (b *BookServiceImpl) Delete(bookId string) error {
	panic("unimplemented")
}

// FindAll implements BookRepository.
func (b *BookServiceImpl) FindAll() (books []model.Book, err error) {
	panic("unimplemented")
}

// FindById implements BookRepository.
func (b *BookServiceImpl) FindById(bookId string) (book model.Book, err error) {
	panic("unimplemented")
}

// Save implements BookRepository.
func (b *BookServiceImpl) Save(book model.Book) (model.Book, error) {
	panic("unimplemented")
}

// Search implements BookRepository.
func (b *BookServiceImpl) Search(s string) (books []model.Book, err error) {
	panic("unimplemented")
}

// Update implements BookRepository.
func (b *BookServiceImpl) Update(book model.Book) (model.Book, error) {
	panic("unimplemented")
}
