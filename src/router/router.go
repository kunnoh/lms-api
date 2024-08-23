package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/data/response"
)

func NewRouter(userCtrl *controller.UserController, authCtrl *controller.AuthController) *gin.Engine {
	routes := gin.Default()

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Message string }{Message: "Welcome to LMS-API"})
	})

	routes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Status string }{Status: "Healthy"})
	})

	authRouter := routes.Group("/auth")
	authRouter.POST("/login", authCtrl.Login)
	authRouter.POST("/register", authCtrl.Register)

	userRouter := routes.Group("/user")
	userRouter.GET("", userCtrl.FindAll)
	userRouter.GET("/:UserId", userCtrl.FindById)
	userRouter.POST("", userCtrl.Create)
	userRouter.PUT("/:UserId", userCtrl.Update)
	userRouter.DELETE("/:UserId", userCtrl.Delete)

	// Catch-all route for handling 404 errors
	routes.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:   404,
			Status: "NotFound",
			Error:  "Page Not Found",
		})
	})

	return routes
}
