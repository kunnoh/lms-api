package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// create controller
func (ctrl *UserController) Create(ctx *gin.Context) {

}

// update controller
func (ctrl *UserController) Update(ctx *gin.Context) {

}

// delete controller
func (ctrl *UserController) Delete(ctx *gin.Context) {

}

// findbyid controller
func (ctrl *UserController) FindById(ctx *gin.Context) {

}

// find all controller
func (ctrl *UserController) FindAll(ctx *gin.Context) {

}
