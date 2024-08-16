package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/controller"
)

func NewRouter(userCtrl *controller.UserController) *gin.Engine {
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Message string }{Message: "Welcome to LMS-API"})
	})

	routes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Message string }{Message: "OK"})
	})

	userRouter := routes.Group("/user")
	userRouter.GET("", userCtrl.FindAll)
	userRouter.GET("/:UserId", userCtrl.FindById)
	userRouter.POST("", userCtrl.Create)
	userRouter.PATCH("/:UserId", userCtrl.Update)
	userRouter.DELETE("", userCtrl.Delete)

	return routes
}
