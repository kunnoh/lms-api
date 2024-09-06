package bookservice

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/model"
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
func (b *BookServiceImpl) Create(book request.CreateBookRequest) response.Response {
	err := b.validate.Struct(book)

	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation failed",
			Error:  err.Error(),
		}
	}

	newBook := model.Book{
		Title:       book.Title,
		ISBN:        book.ISBN,
		Publication: book.Publication,
		Author:      book.Author,
		Genre:       book.Genre,
	}

	savedBk, err := b.BookRepo.Save(newBook)

	if err != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error saving the book",
			Error:  err.Error(),
		}
	}

	return response.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data: response.BookResponse{
			BookId:      savedBk.BookId,
			Title:       savedBk.Title,
			ISBN:        savedBk.ISBN,
			Publication: savedBk.Publication,
			Author:      savedBk.Author,
			Genre:       savedBk.Genre,
		},
	}
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(BookId string) response.Response {
	panic("unimplemented")
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll() response.Response {
	res, err := b.BookRepo.FindAll()
	if err != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "error",
			Error:  "Failed to fetch books",
		}
	}

	books := make([]response.BookResponse, 0, len(res))

	for _, val := range res {
		books = append(books, response.BookResponse{
			BookId: val.BookId,
			Title:  val.Title,
		})
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   books,
	}

}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(BookId string) response.Response {
	bookData, err := b.BookRepo.FindById(BookId)

	if err != nil {
		return response.Response{
			Code:   http.StatusNotFound,
			Status: "Error",
			Error:  "Book not found",
		}
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data: response.BookResponse{
			BookId:      bookData.BookId,
			Title:       bookData.Title,
			ISBN:        bookData.ISBN,
			Publication: bookData.Publication,
			Genre:       bookData.Genre,
			Author:      bookData.Author,
		},
	}
}

// Update implements BookService.
func (b *BookServiceImpl) Update() response.Response {
	panic("unimplemented")
}
