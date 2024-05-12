package mocks

import (
	"mini-project/transaction"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) GetTransactionByOrderId(id transaction.GetOrderTransactionInput) ([]transaction.Transaction, error) {

	args := u.Called(id)

	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (u *Usecase) GetTransactionByUserId(id int) ([]transaction.Transaction, error) {

	args := u.Called(id)

	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (u *Usecase) CreateTransaction(input transaction.CreateTransactionInput) (transaction.Transaction, error) {

	args := u.Called(input)

	return args.Get(0).(transaction.Transaction), args.Error(1)
}

func (u *Usecase) PaymentProcess(input transaction.TransactionNotificationInput) error {

	args := u.Called(input)

	return args.Error(0)
}

func NewUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}
