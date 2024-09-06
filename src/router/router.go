package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/middleware"
	userrepository "github.com/kunnoh/lms-api/src/repository/user.repository"
)

func NewRouter(usersRepo userrepository.UserRepository, userCtrl *controller.UserController, bookCtrl *controller.BookController, authCtrl *controller.AuthController) *gin.Engine {
	routes := gin.Default()

	routes.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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

	bookRouter := routes.Group("/books")
	// bookRouter.GET("", bookCtrl.FindAll)
	bookRouter.GET("", middleware.DeserializeUser(usersRepo), bookCtrl.FindAll)
	bookRouter.GET("/:BookId", middleware.DeserializeUser(usersRepo), bookCtrl.FindById)
	bookRouter.POST("", middleware.DeserializeUser(usersRepo), bookCtrl.Create)
	bookRouter.PUT("/:BookId", middleware.DeserializeUser(usersRepo), bookCtrl.Update)
	bookRouter.DELETE("/:BookId", middleware.DeserializeUser(usersRepo), bookCtrl.Delete)

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
