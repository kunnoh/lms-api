package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/middleware"
	"github.com/kunnoh/lms-api/src/repository"
)

func NewRouter(usersRepo repository.UserRepository, userCtrl *controller.UserController, authCtrl *controller.AuthController) *gin.Engine {
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
	authRouter.GET("/refresh-token", authCtrl.RefreshToken)

	userRouter := routes.Group("/user")
	userRouter.GET("", middleware.DeserializeUser(usersRepo), userCtrl.FindAll)
	userRouter.GET("/:UserId", middleware.DeserializeUser(usersRepo), userCtrl.FindById)
	userRouter.POST("", middleware.DeserializeUser(usersRepo), userCtrl.Create)
	userRouter.PUT("/:UserId", middleware.DeserializeUser(usersRepo), userCtrl.Update)
	userRouter.DELETE("/:UserId", middleware.DeserializeUser(usersRepo), userCtrl.Delete)

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
