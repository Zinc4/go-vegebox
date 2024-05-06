package handler

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/product"
	"mini-project/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	usecase     admin.Usecase
	authUsecase auth.Usecase
}

func NewAdminHandler(usecase admin.Usecase, authUsecase auth.Usecase) *adminHandler {
	return &adminHandler{usecase, authUsecase}
}

func (h *adminHandler) CreateProduct(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}

	var input product.AddProductInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newProduct, err := h.usecase.CreateProduct(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	formatter := product.FormatProduct(newProduct)

	return c.JSON(http.StatusOK, formatter)
}

func (h *adminHandler) CreateCategory(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}

	var input product.Category
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	category, err := h.usecase.CreateCategory(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, category)

}
