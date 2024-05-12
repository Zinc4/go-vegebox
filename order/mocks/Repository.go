package mocks

import (
	"mini-project/order"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Create(odr order.Order) (*order.Order, error) {

	args := r.Called(odr)
	return args.Get(0).(*order.Order), args.Error(1)
}

func (r *Repository) Update(odr order.Order) (order.Order, error) {

	args := r.Called(odr)
	return *args.Get(0).(*order.Order), args.Error(1)
}

func (r *Repository) GetOrderByID(id uuid.UUID) (order.Order, error) {

	args := r.Called(id)
	return *args.Get(0).(*order.Order), args.Error(1)
}

func (r *Repository) CreateOrderItems(orderItems []order.OrderItem) error {

	args := r.Called(orderItems)
	return args.Error(0)
}

func (r *Repository) GetOrderItemsByOrderID(orderID uuid.UUID) ([]order.OrderItem, error) {

	args := r.Called(orderID)
	return args.Get(0).([]order.OrderItem), args.Error(1)
}

func (r *Repository) GetOrdersByUserID(userID int) ([]order.Order, error) {

	args := r.Called(userID)
	return args.Get(0).([]order.Order), args.Error(1)
}

func (r *Repository) GetOrderItemByID(id uuid.UUID) (order.Order, error) {

	args := r.Called(id)
	return args.Get(0).(order.Order), args.Error(1)
}

func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
