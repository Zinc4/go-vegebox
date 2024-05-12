package mocks

import (
	"errors"
	"mini-project/order"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = order.NewUsecase(repository)

	var order = order.Order{
		ID:        uuid.New(),
		UserID:    1,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	t.Run("Success create order", func(t *testing.T) {

		repository.On("Create", order).Return(&order, nil).Once()

		result, err := usecase.CreateOrder(order)

		assert.NoError(t, err)
		assert.Equal(t, &order, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed create order", func(t *testing.T) {

		repository.On("Create", order).Return(&order, errors.New("error")).Once()

		result, err := usecase.CreateOrder(order)

		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		assert.Equal(t, &order, result)
		repository.AssertExpectations(t)
	})

}

func TestCreateOrderItems(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = order.NewUsecase(repository)

	var orderItems = []order.OrderItem{
		{

			OrderID:   uuid.New(),
			ProductID: 1,
			Quantity:  1,
		},
		{

			OrderID:   uuid.New(),
			ProductID: 1,
			Quantity:  1,
		},
	}

	t.Run("Success create order items", func(t *testing.T) {

		repository.On("CreateOrderItems", orderItems).Return(nil).Once()

		err := usecase.CreateOrderItems(orderItems)

		assert.NoError(t, err)

		repository.AssertExpectations(t)
	})

	t.Run("Failed create order items", func(t *testing.T) {

		repository.On("CreateOrderItems", orderItems).Return(errors.New("error")).Once()

		err := usecase.CreateOrderItems(orderItems)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		repository.AssertExpectations(t)
	})

}

func TestGetOrderItemsByOrderID(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = order.NewUsecase(repository)

	var orderItems = []order.OrderItem{
		{
			OrderID:   uuid.New(),
			ProductID: 1,
			Quantity:  1,
		},
		{
			OrderID:   uuid.New(),
			ProductID: 1,
			Quantity:  1,
		},
	}

	t.Run("Success get order items by order id", func(t *testing.T) {

		repository.On("GetOrderItemsByOrderID", orderItems[0].OrderID).Return(orderItems, nil).Once()

		result, err := usecase.GetOrderItemsByOrderID(orderItems[0].OrderID)

		assert.NoError(t, err)

		assert.Equal(t, orderItems, result)

		repository.AssertExpectations(t)
	})

	t.Run("Failed get order items by order id", func(t *testing.T) {

		repository.On("GetOrderItemsByOrderID", orderItems[0].OrderID).Return(orderItems, errors.New("error")).Once()

		result, err := usecase.GetOrderItemsByOrderID(orderItems[0].OrderID)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		assert.Equal(t, orderItems, result)

		repository.AssertExpectations(t)
	})
}

func TestGetOrderItemByID(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = order.NewUsecase(repository)

	var order = order.Order{
		ID:        uuid.New(),
		UserID:    1,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	t.Run("Success get order item by id", func(t *testing.T) {

		repository.On("GetOrderItemByID", order.ID).Return(order, nil).Once()

		result, err := usecase.GetOrderItemByID(order.ID)

		assert.NoError(t, err)

		assert.Equal(t, order, result)

		repository.AssertExpectations(t)
	})

	t.Run("Failed get order item by id", func(t *testing.T) {

		repository.On("GetOrderItemByID", order.ID).Return(order, errors.New("error")).Once()

		result, err := usecase.GetOrderItemByID(order.ID)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		assert.Equal(t, order, result)

		repository.AssertExpectations(t)
	})
}

func TestGetOrdersByUserID(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = order.NewUsecase(repository)

	var odr = order.Order{
		ID:        uuid.New(),
		UserID:    1,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	t.Run("Success get orders by user id", func(t *testing.T) {

		repository.On("GetOrdersByUserID", odr.UserID).Return([]order.Order{odr}, nil).Once()

		result, err := usecase.GetOrdersByUserID(odr.UserID)

		assert.NoError(t, err)

		assert.Equal(t, []order.Order{odr}, result)

		repository.AssertExpectations(t)
	})

	t.Run("Failed get orders by user id", func(t *testing.T) {

		repository.On("GetOrdersByUserID", odr.UserID).Return([]order.Order{odr}, errors.New("error")).Once()

		result, err := usecase.GetOrdersByUserID(odr.UserID)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		assert.Equal(t, []order.Order{odr}, result)

		repository.AssertExpectations(t)

	})

}
