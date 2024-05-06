package handler

import (
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase user.Usecase
	authUsecase auth.Usecase
}

func NewUserHandler(userUsecase user.Usecase, authUsecase auth.Usecase) *userHandler {
	return &userHandler{userUsecase, authUsecase}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	input := user.RegisterUserInput{}
	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		response := helper.ErrorResponse("failed to register account", errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = h.userUsecase.RegisterUser(input)
	if err != nil {
		response := helper.ErrorResponse("failed to register account", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("success register account, please check your email for verification"))
}
