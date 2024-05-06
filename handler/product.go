package handler

import (
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	formatter := product.FormatProducts(products)

	return c.JSON(http.StatusOK, formatter)
}

func (h *productHandler) GetProductByID(c echo.Context) error {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := h.productUsecase.FindProductByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}
