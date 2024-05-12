package handler

import (
	"mini-project/helper"
	"mini-project/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUsecase product.Usecase
}

func NewProductHandler(productUsecase product.Usecase) *productHandler {
	return &productHandler{productUsecase}
}

func (h *productHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.FindProducts()
	if err != nil {
		response := helper.ErrorResponse("Error get products", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := product.FormatProducts(products)

	return c.JSON(http.StatusOK, helper.ResponseWithData("success get products", formatter))
}

func (h *productHandler) GetProductByID(c echo.Context) error {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := h.productUsecase.FindProductByID(id)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := product.FormatProduct(result)

	return c.JSON(http.StatusOK, helper.ResponseWithData("success get product", formatter))
}

func (h *productHandler) GetAllCategory(c echo.Context) error {

	categories, err := h.productUsecase.FindAllCategory()
	if err != nil {
		response := helper.ErrorResponse("Error get categories", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData("success get categories", categories))
}

func (h *productHandler) GetProductByCategory(c echo.Context) error {

	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	products, err := h.productUsecase.FindProductByCategory(id)
	if err != nil {
		response := helper.ErrorResponse("Error get products", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := product.FormatProducts(products)

	return c.JSON(http.StatusOK, helper.ResponseWithData("success get products", formatter))
}
