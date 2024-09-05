package bookservice

import (
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
)

type BookService interface {
	Create(Book request.CreateBookRequest) response.Response
	Update() response.Response
	Delete(BookId string) response.Response
	FindById(BookId string) response.Response
	FindAll() response.Response
}
