package transaction

import (
	"mini-project/helper"
	"mini-project/order"
	"mini-project/payment"
	"strconv"
)

type Usecase interface {
	GetTransactionByUserId(id int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	PaymentProcess(input TransactionNotificationInput) error
}

type usecase struct {
	repository      Repository
	orderRepository order.Repository
	paymentUsecase  payment.Usecase
}

func NewUsecase(repository Repository, orderRepository order.Repository, paymentUsecase payment.Usecase) *usecase {
	return &usecase{repository, orderRepository, paymentUsecase}
}

func (u *usecase) GetTransactionByUserId(id int) ([]Transaction, error) {
	transactions, err := u.repository.GetByUserId(id)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (u *usecase) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.OrderID = input.OrderID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	transaction.Amount = input.Amount

	Random := helper.GenerateRandomOTP(5)
	atoi, err := strconv.Atoi(Random)
	if err != nil {
		return Transaction{}, err
	}

	paymenyTransaction := payment.Transaction{
		ID:     atoi,
		Amount: input.Amount,
	}

	paymentUrl, err := u.paymentUsecase.GetPaymentURL(paymenyTransaction, input.User)
	if err != nil {
		return transaction, err
	}

	transaction.PaymentURL = paymentUrl
	transaction.Code = strconv.Itoa(paymenyTransaction.ID)

	newTransaction, err := u.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}

func (u *usecase) PaymentProcess(input TransactionNotificationInput) error {
	code := input.OrderID

	transaction, err := u.repository.GetByCode(code)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" || input.TransactionStatus == "capture" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "cancel" || input.TransactionStatus == "expired" {
		transaction.Status = "failed"
	} else if input.TransactionStatus == "pending" || input.TransactionStatus == "challenge" {
		transaction.Status = "pending"
	}

	updatedTransaction, err := u.repository.Update(transaction)
	if err != nil {
		return err
	}

	order, err := u.orderRepository.GetOrderByID(updatedTransaction.OrderID)
	if err != nil {
		return err
	}

	orderItem, err := u.orderRepository.GetOrderItemsByOrderID(updatedTransaction.OrderID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		for _, item := range orderItem {
			item.Order.Status = "paid"
		}
		order.Status = "paid"
		_, err = u.orderRepository.Update(order)
		if err != nil {
			return err
		}
	}

	return nil

}
