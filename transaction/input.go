package transaction

import (
	"mini-project/user"

	"github.com/google/uuid"
)

type GetOrderTransactionInput struct {
	ID   uuid.UUID `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	OrderID uuid.UUID `json:"order_id" binding:"required"`
	User    user.User
	Amount  int `json:"amount" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
