package bookrepository

import (
	"errors"
	"fmt"

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
	res := b.Db.Where("book_id = ?", bookId).Delete(&model.Book{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", bookId)
	}

	return nil
}

// FindAll implements BookRepository.
func (b *BookServiceImpl) FindAll() ([]model.Book, error) {
	var bks []model.Book

	res := b.Db.Find(&bks)

	if res.Error != nil {
		return nil, res.Error
	}

	return bks, nil
}

// FindById implements BookRepository.
func (b *BookServiceImpl) FindById(bookId string) (model.Book, error) {
	var bk model.Book

	res := b.Db.First(&bk, "book_id = ?", bookId)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return bk, errors.New("book not found")
		}
		return bk, res.Error
	}
	return bk, nil
}

// Save implements BookRepository.
func (b *BookServiceImpl) Save(book model.Book) (model.Book, error) {
	res := b.Db.Create(&book)

	if res.Error != nil {
		return book, res.Error
	}
	return book, nil
}

// Search implements BookRepository.
func (b *BookServiceImpl) Search(s string) ([]model.Book, error) {
	panic("unimplemented")
}

// Update implements BookRepository.
func (b *BookServiceImpl) Update(book model.Book) (model.Book, error) {
	panic("unimplemented")

}
