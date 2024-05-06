package routes

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/handler"
	"mini-project/middleware"
	"mini-project/user"
	"mini-project/utils/database"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {
	userRepository := user.NewRepository(database.DB)
	adminRepository := admin.NewRepository(database.DB)

	authUsecase := auth.NewUsecase()
	userUsecase := user.NewUsecase(userRepository)
	adminUsecase := admin.NewUsecase(adminRepository)

	userHandler := handler.NewUserHandler(userUsecase, authUsecase)
	adminHandler := handler.NewAdminHandler(adminUsecase, authUsecase)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/verify-email", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)

	api.POST("/products", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.CreateProduct))

}
