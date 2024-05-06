package handler

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/product"
	"mini-project/user"
	"net/http"
	"strconv"

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

func (h *adminHandler) GetAllUsers(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}

	pageStr := c.QueryParam("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := c.QueryParam("page_size")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)

	var users []user.User
	var totalPages, currentPage, nextPage, prevPage int
	var err error

	searchNameUser := c.QueryParam("name")
	if searchNameUser != "" {
		users, err = h.usecase.SearchUserByName(searchNameUser)
	} else {
		users, totalPages, currentPage, nextPage, prevPage, err = h.usecase.GetUserPagination(page, pageSize)
	}

	if err != nil {
		respone := helper.ErrorResponse("Failed to get users", err.Error())
		return c.JSON(http.StatusBadRequest, respone)
	}

	var nonAdminUsers []user.User
	for _, user := range users {
		if user.Role != "admin" {
			nonAdminUsers = append(nonAdminUsers, user)
		}
	}

	if totalPages > 1 {
		if currentPage < totalPages {
			nextPage = currentPage + 1
		} else {
			nextPage = -1
		}
		if currentPage > 1 {
			prevPage = currentPage - 1
		} else {
			prevPage = -1
		}
	}

	response := helper.ResponseWithPaginationAndNextPrev("Get users success", admin.FormatterUsers(nonAdminUsers), currentPage, totalPages, nextPage, prevPage)

	return c.JSON(http.StatusOK, response)

}
