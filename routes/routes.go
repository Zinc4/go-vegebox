package routes

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/handler"
	"mini-project/middleware"
	"mini-project/product"
	"mini-project/user"
	"mini-project/utils/database"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {
	userRepository := user.NewRepository(database.DB)
	adminRepository := admin.NewRepository(database.DB)

	productRepository := product.NewRepository(database.DB)

	productUsecase := product.NewUsecase(productRepository)

	authUsecase := auth.NewUsecase()
	userUsecase := user.NewUsecase(userRepository)
	adminUsecase := admin.NewUsecase(adminRepository)

	productHandler := handler.NewProductHandler(productUsecase)

	userHandler := handler.NewUserHandler(userUsecase, authUsecase)
	adminHandler := handler.NewAdminHandler(adminUsecase, authUsecase)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/verify-email", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)

	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:id", productHandler.GetProductByID)

	api.POST("/products", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.CreateProduct))
	api.POST("/category", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.CreateCategory))
	api.GET("/users", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.GetAllUsers))

}
