package routes

import (
	"mini-project/auth"
	"mini-project/handler"
	"mini-project/user"
	"mini-project/utils/database"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {
	userRepository := user.NewRepository(database.DB)

	authUsecase := auth.NewUsecase()
	userUsecase := user.NewUsecase(userRepository)

	userHandler := handler.NewUserHandler(userUsecase, authUsecase)

	api := router.Group("api/v1")

	api.POST("/register", userHandler.RegisterUser)

}
