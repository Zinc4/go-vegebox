package mocks

import (
	"mini-project/order"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) CreateOrder(odr order.Order) (*order.Order, error) {

	args := u.Called(odr)
	return args.Get(0).(*order.Order), args.Error(1)
}

func (u *Usecase) CreateOrderItems(orderItems []order.OrderItem) error {

	args := u.Called(orderItems)
	return args.Error(0)
}

func (u *Usecase) GetOrderItemsByOrderID(orderID uuid.UUID) ([]order.OrderItem, error) {

	args := u.Called(orderID)
	return args.Get(0).([]order.OrderItem), args.Error(1)
}

func (u *Usecase) GetOrderItemByID(id uuid.UUID) (order.Order, error) {

	args := u.Called(id)
	return *args.Get(0).(*order.Order), args.Error(1)
}

func (u *Usecase) GetOrdersByUserID(userID int) ([]order.Order, error) {

	args := u.Called(userID)
	return args.Get(0).([]order.Order), args.Error(1)
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
