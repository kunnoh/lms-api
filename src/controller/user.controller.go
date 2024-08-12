package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/services"
	"github.com/kunnoh/lms-api/src/utils"
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
	createUserReq := request.CreateUserRequest{}
	// err := ctx.ShouldBindJSON(&createUserReq)
	// utils.ErrorPanic(err)

	if err := ctx.ShouldBindJSON(&createUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}

	ctrl.userService.Create(createUserReq)

	webResp := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
	}
	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResp)
}

// update controller
func (ctrl *UserController) Update(ctx *gin.Context) {
	updateUserReq := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserReq)
	utils.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)
	updateUserReq.UserId = id

	ctrl.userService.Update(updateUserReq)

	webResp := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResp)
}

// delete controller
func (ctrl *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)
	ctrl.userService.Delete(id)

	webResp := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResp)
}

// findbyid controller
func (ctrl *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)

	userResp := ctrl.userService.FindById(id)

	webResp := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResp,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResp)
}

// find all controller
func (ctrl *UserController) FindAll(ctx *gin.Context) {
	userResp := ctrl.userService.FindAll()

	webResp := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResp,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResp)
}
