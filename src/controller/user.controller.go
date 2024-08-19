package controller

import (
	// "fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
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

// Create controller
func (ctrl *UserController) Create(ctx *gin.Context) {
	var createUserReq request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&createUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}

	res := ctrl.userService.Create(createUserReq)
	ctx.JSON(res.Code, res)
}

// Update controller
func (ctrl *UserController) Update(ctx *gin.Context) {
	var updateUserReq request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&updateUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}

	userId := ctx.Param("userId")
	// id, err := strconv.Atoi(userId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.Response{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Error:  "Invalid user ID",
	// 	})
	// 	return
	// }
	updateUserReq.UserId = userId

	ctrl.userService.Update(updateUserReq)
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// Delete controller
func (ctrl *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	// id, err := strconv.Atoi(userId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.Response{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Error:  "Invalid user ID",
	// 	})
	// 	return
	// }

	ctrl.userService.Delete(userId)
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// FindById controller
func (ctrl *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("UserId")
	// id, err := strconv.Atoi(userId)
	// fmt.Println(err)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.Response{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Error:  "Invalid user ID",
	// 	})
	// 	return
	// }

	userResp := ctrl.userService.FindById(userId)
	ctx.JSON(userResp.Code, userResp) // Return the response from the service directly
}

// FindAll controller
func (ctrl *UserController) FindAll(ctx *gin.Context) {
	userResp := ctrl.userService.FindAll()
	ctx.JSON(userResp.Code, userResp) // Return the response from the service directly
}
