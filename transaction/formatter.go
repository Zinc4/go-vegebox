package transaction

import (
	"time"

	"github.com/google/uuid"
)

type OrderTransactionFormatter struct {
	ID         int       `json:"id"`
	OrderID    uuid.UUID `json:"order_id"`
	UserID     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
}

func FormatOrderTransaction(transaction Transaction) OrderTransactionFormatter {
	formatter := OrderTransactionFormatter{
		ID:         transaction.ID,
		OrderID:    transaction.OrderID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}

	return formatter
}

func FormatOrderTransactions(transactions []Transaction) []OrderTransactionFormatter {
	if len(transactions) == 0 {
		return []OrderTransactionFormatter{}
	}
	var orderTransactions []OrderTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatOrderTransaction(transaction)
		orderTransactions = append(orderTransactions, formatter)
	}

	return orderTransactions
}

type OrderFormatter struct {
	OrderID uuid.UUID   `json:"order_id"`
	Item    interface{} `json:"item"`
}

type UserTransactionFormatter struct {
	ID        int            `json:"id"`
	Amount    int            `json:"amount"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	Order     OrderFormatter `json:"order"`
}

func FormatUserTransaction(transaction Transaction, orderItem interface{}) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	formatOrder := OrderFormatter{}

	formatOrder.OrderID = transaction.OrderID
	formatOrder.Item = orderItem

	formatter.Order = formatOrder

	return formatter

}

// func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
// 	if len(transactions) == 0 {
// 		return []UserTransactionFormatter{}
// 	}

// 	var transactionsFormatter []UserTransactionFormatter

// 	for _, transaction := range transactions {
// 		formatter := FormatUserTransaction(transaction)
// 		transactionsFormatter = append(transactionsFormatter, formatter)
// 	}

// 	return transactionsFormatter
// }
