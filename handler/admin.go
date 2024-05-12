package handler

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/product"
	"mini-project/transaction"
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

func (h *adminHandler) DeleteUser(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.ErrorResponse("Failed get user", err.Error())
		return c.JSON(http.StatusBadRequest, response)

	}
	user, err := h.usecase.FindUserById(userId)
	if err != nil {
		response := helper.ErrorResponse("Failed get user", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.usecase.DeleteUserById(user.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed to delete user", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := helper.SuccesResponse("Delete user success")
	return c.JSON(http.StatusOK, response)
}

func (h *adminHandler) GetAllUserTransactions(c echo.Context) error {
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
		pageSizeStr = "20"
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)

	var transactions []transaction.Transaction
	var totalPages, currentPage, nextPage, prevPage int
	var err error

	searchTransactionByName := c.QueryParam("name")
	if searchTransactionByName != "" {
		transactions, err = h.usecase.SearchTransactionByName(searchTransactionByName)
	} else {
		transactions, totalPages, currentPage, nextPage, prevPage, err = h.usecase.GetTransactionsPagination(page, pageSize)
	}

	if err != nil {
		respone := helper.ErrorResponse("Failed to get transactions", err.Error())
		return c.JSON(http.StatusBadRequest, respone)
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

	response := helper.ResponseWithPaginationAndNextPrev("Get transactions success", admin.FormatTransactions(transactions), currentPage, totalPages, nextPage, prevPage)

	return c.JSON(http.StatusOK, response)
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
		response := helper.ErrorResponse("Failed to create product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	newProduct, err := h.usecase.CreateProduct(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to create product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := product.FormatProduct(newProduct)

	return c.JSON(http.StatusOK, helper.ResponseWithData("Create product success", formatter))
}

func (h *adminHandler) UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := helper.ErrorResponse("Failed get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	var input product.GetProductDetailInput
	input.ID = id

	var data product.AddProductInput
	err = c.Bind(&data)
	if err != nil {
		response := helper.ErrorResponse("Failed to update product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedProduct, err := h.usecase.UpdateProduct(input, data)
	if err != nil {
		response := helper.ErrorResponse("Failed to update product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData("Update product success", updatedProduct))
}

func (h *adminHandler) DeleteProduct(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}

	product, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.ErrorResponse("Failed get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	productId, err := h.usecase.FindProductByID(product)
	if err != nil {
		response := helper.ErrorResponse("Failed to get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.usecase.DeleteProductByID(productId.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed to delete product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := helper.GeneralResponse("Delete product success")
	return c.JSON(http.StatusOK, response)

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
		response := helper.ErrorResponse("Failed to create category", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.usecase.CreateCategory(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to create category", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData("Create category success", category))

}

func (h *adminHandler) DeleteCategory(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		respone := helper.GeneralResponse("Access Denied")
		return c.JSON(http.StatusUnauthorized, respone)
	}

	category, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.ErrorResponse("Failed get category", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	categoryId, err := h.usecase.FindCategoryByID(category)
	if err != nil {
		response := helper.ErrorResponse("Failed to get category", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.usecase.DeleteCategoryByID(categoryId.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed to delete category", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.GeneralResponse("Delete category success")

	return c.JSON(http.StatusOK, response)
}
