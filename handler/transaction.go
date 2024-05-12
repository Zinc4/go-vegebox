package handler

import (
	"mini-project/helper"
	"mini-project/order"
	"mini-project/transaction"
	"mini-project/user"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUsecase transaction.Usecase
	orderUsecase       order.Usecase
}

func NewTransactionHandler(transactionUsecase transaction.Usecase, orderUsecase order.Usecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase, orderUsecase}
}

func (h *TransactionHandler) GetUserTransaction(c echo.Context) error {

	currentUser := c.Get("CurrentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.transactionUsecase.GetTransactionByUserId(userID)
	if err != nil {
		response := helper.ErrorResponse("Failed to get user's transactions", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	var formattedTransactions []transaction.UserTransactionFormatter
	for _, transactionn := range transactions {
		items, _ := h.orderUsecase.GetOrderItemsByOrderID(transactionn.Order.ID)
		orderItems := make([]order.ResponseOrderItem, len(items))
		for i, item := range items {
			orderItems[i].Product = item.Product
			orderItems[i].Quantity = item.Quantity
			orderItems[i].Price = item.Product.Price * item.Quantity
		}

		formattedTransactions = append(formattedTransactions, transaction.FormatUserTransaction(transactionn, orderItems))

	}

	response := helper.ResponseWithData("Success to get user's transactions", formattedTransactions)
	return c.JSON(http.StatusOK, response)
}

func calculateTotalOrderAmount(orderItems []order.OrderItem) int {
	total := 0
	for _, item := range orderItems {
		total += item.Quantity * item.Price
	}
	return total
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	user := c.Get("CurrentUser").(user.User)

	var input transaction.CreateTransactionInput

	id := c.Param("order_id")
	orderId, err := uuid.Parse(id)
	if err != nil {
		response := helper.ErrorResponse("Error get order", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	input.OrderID = orderId
	order, err := h.orderUsecase.GetOrderItemsByOrderID(input.OrderID)
	if err != nil {
		response := helper.ErrorResponse("Error get order", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	totalOrderAmount := calculateTotalOrderAmount(order)

	input.Amount = totalOrderAmount
	input.User = user
	input.OrderID = orderId

	newTransaction, err := h.transactionUsecase.CreateTransaction(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to create transaction", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success to create transaction", transaction.FormatOrderTransaction(newTransaction))
	return c.JSON(http.StatusOK, response)

}

func (h *TransactionHandler) GetPaymentCallback(c echo.Context) error {

	var input transaction.TransactionNotificationInput
	err := c.Bind(&input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get payment callback", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.transactionUsecase.PaymentProcess(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get payment callback", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, input)

}
