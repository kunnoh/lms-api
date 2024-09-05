package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
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
	var creteBookReq request.CreateBookRequest
	if err := ctx.ShouldBindJSON(&creteBookReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}

	bkRes := ctrl.bookService.Create(creteBookReq)
	ctx.JSON(bkRes.Code, bkRes)
}

// FindById controller
func (ctrl *BookController) FindById(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	bkResp := ctrl.bookService.FindById(bookId)
	ctx.JSON(bkResp.Code, bkResp)
}

// Update controller
func (ctrl *BookController) Update(ctx *gin.Context) {
	// bookId := ctx.Param("BookId")
	// bkResp := ctrl.bookService.Update(bookId)
	// ctx.JSON(bkResp.Code, bkResp)
}

// FindAll controller
func (ctrl *BookController) FindAll(ctx *gin.Context) {
	bkResp := ctrl.bookService.FindAll()
	ctx.JSON(bkResp.Code, bkResp)
}

// Delete controller
func (ctrl *BookController) Delete(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	delResp := ctrl.bookService.Delete(bookId)
	ctx.JSON(delResp.Code, delResp)
}
