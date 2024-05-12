package admin

import (
	"mini-project/transaction"
	"mini-project/user"
	"time"

	"github.com/google/uuid"
)

type UserFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatUser(user user.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Avatar = user.Avatar
	userFormatter.Role = user.Role
	userFormatter.IsVerified = user.IsVerified
	userFormatter.CreatedAt = user.CreateAt

	return userFormatter
}

func FormatterUsers(users []user.User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		formatUser := FormatUser(user)
		usersFormatter = append(usersFormatter, formatUser)
	}

	return usersFormatter
}

type TransactionFormatter struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	OrderID    uuid.UUID `json:"order_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
	User       string    `json:"user"`
	CreateAt   time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatTransaction(transaction transaction.Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.UserID = transaction.UserID
	transactionFormatter.OrderID = transaction.OrderID
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.Status = transaction.Status
	transactionFormatter.Code = transaction.Code
	transactionFormatter.PaymentURL = transaction.PaymentURL
	transactionFormatter.User = transaction.User.Name
	transactionFormatter.CreateAt = transaction.CreatedAt
	transactionFormatter.UpdatedAt = transaction.UpdatedAt

	return transactionFormatter

}

func FormatTransactions(transactions []transaction.Transaction) []TransactionFormatter {
	if len(transactions) == 0 {
		return []TransactionFormatter{}
	}
	var transactionsFormatter []TransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
