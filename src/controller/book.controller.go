package controller

import (
	"github.com/gin-gonic/gin"
	bookservice "github.com/kunnoh/lms-api/src/services/book.service"
)

type BookController struct {
	bookService bookservice.BookService
}

func NewBookController(service bookservice.BookService) *BookController {
	return &BookController{
		bookService: service,
	}
}

// Create controller
func (ctrl *BookController) Create(ctx *gin.Context) {

}

// FindById controller
func (ctrl *BookController) FindById(ctx *gin.Context) {

}

// Update controller
func (ctrl *BookController) Update(ctx *gin.Context) {

}

// FindAll controller
func (ctrl *BookController) FindAll(ctx *gin.Context) {
	userResp := ctrl.bookService.FindAll()
	ctx.JSON(userResp.Code, userResp)
}

// Delete controller
func (ctrl *BookController) Delete(ctx *gin.Context) {

}
