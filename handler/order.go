package handler

import (
	"mini-project/cart"
	"mini-project/helper"
	"mini-project/order"
	"mini-project/product"
	"mini-project/user"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase   order.Usecase
	cartUsecase    cart.Usecase
	productUsecase product.Usecase
}

func NewOrderHandler(orderUsecase order.Usecase, cartUsecase cart.Usecase, productUsecase product.Usecase) *OrderHandler {
	return &OrderHandler{orderUsecase, cartUsecase, productUsecase}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	user, _ := c.Get("CurrentUser").(user.User)

	id := c.Param("cart_id")
	cartId, err := strconv.Atoi(id)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cart, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	items, _ := h.cartUsecase.GetCartItemByCartId(cart.ID)
	if len(items) == 0 {
		response := helper.ErrorResponse("Cart is empty", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	newOrder := order.Order{
		ID:     uuid.New(),
		UserID: user.ID,
		Status: "unpaid",
	}

	orderItems := make([]order.OrderItem, len(items))
	for i, item := range items {

		orderItems[i] = order.OrderItem{
			OrderID:   newOrder.ID,
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}

	}

	order_res, err := h.orderUsecase.CreateOrder(newOrder)
	if err != nil {
		response := helper.ErrorResponse("Error create order", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.orderUsecase.CreateOrderItems(orderItems)
	if err != nil {
		response := helper.ErrorResponse("Error create order items", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, item := range items {
		err = h.cartUsecase.DeleteCartItemByCartIdAndProductId(cart.ID, item.ProductID)
		if err != nil {
			return err
		}
	}

	err = h.cartUsecase.DeleteCartById(cart.ID)
	if err != nil {
		return err
	}

	response := helper.ResponseWithData("Success create order", h.orderUsecase.GetOrderSerializer(*order_res))

	return c.JSON(http.StatusOK, response)

}

func (h *OrderHandler) GetOrder(c echo.Context) error {

	id := c.Param("order_id")

	orderId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid order id")
	}

	orderItems, err := h.orderUsecase.GetOrderItemsByOrderID(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	formatter := order.FormatOrders(orderItems)

	return c.JSON(http.StatusOK, helper.ResponseWithData("Success get order", formatter))
}

func (h *OrderHandler) GetOrders(c echo.Context) error {

	user, _ := c.Get("CurrentUser").(user.User)
	response := helper.ResponseWithData("Success get orders", h.orderUsecase.GetOrdersSerializerByUserID(user))
	return c.JSON(http.StatusOK, response)
}
