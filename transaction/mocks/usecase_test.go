package mocks

import (
	"errors"
	"mini-project/order/mocks"
	"mini-project/payment"
	"mini-project/transaction"
	"mini-project/user"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// func TestGetTransactionByOrderId(t *testing.T) {

// 	var repository = NewRepository(t)
// 	var orderRepository = mocks.NewRepository(t)
// 	var paymentUsecase = payment.NewUsecase()
// 	var usecase = transaction.NewUsecase(repository, orderRepository, paymentUsecase)

// 	var orderr = order.Order{
// 		ID:        uuid.New(),
// 		UserID:    1,
// 		Status:    "pending",
// 		CreatedAt: time.Now(),
// 	}

// 	var orderItems = []order.OrderItem{
// 		{
// 			OrderID:   uuid.New(),
// 			ProductID: 1,
// 			Quantity:  1,
// 		},
// 		{
// 			OrderID:   uuid.New(),
// 			ProductID: 2,
// 			Quantity:  2,
// 		},
// 	}

// 	t.Run("Success get transaction by order id", func(t *testing.T) {

// 		orderRepository.On("GetOrderItemByID", orderr.ID).Return(orderr, nil).Once()

// 		orderRepository.On("GetOrderItemsByOrderID", orderr.ID).Return(orderItems, nil).Once()

// 		result, err := usecase.GetTransactionByOrderId(transaction.GetOrderTransactionInput{

// 			ID: orderr.ID,
// 		})

// 		assert.NoError(t, err)
// 		assert.Equal(t, []transaction.Transaction{}, result)

// 		orderRepository.AssertExpectations(t)
// 	})

// 	t.Run("Failed get transaction by order id", func(t *testing.T) {

// 		orderRepository.On("GetOrderItemByID", orderr.ID).Return(orderr, nil).Once()

// 		orderRepository.On("GetOrderItemsByOrderID", orderr.ID).Return(orderItems, errors.New("error")).Once()

// 		result, err := usecase.GetTransactionByOrderId(transaction.GetOrderTransactionInput{

// 			ID: orderr.ID,
// 		})

// 		assert.Error(t, err)

// 		assert.Equal(t, []transaction.Transaction{}, result)

// 		orderRepository.AssertExpectations(t)
// 	})
// }

func TestGetTransactionByUserId(t *testing.T) {

	var repository = NewRepository(t)
	var orderRepository = mocks.NewRepository(t)
	var paymentUsecase = payment.NewUsecase()
	var usecase = transaction.NewUsecase(repository, orderRepository, paymentUsecase)

	userData := user.User{
		ID: 1,
	}

	var transactions = []transaction.Transaction{
		{
			ID:         1,
			OrderID:    uuid.New(),
			UserID:     1,
			Status:     "paid",
			Code:       "12345",
			Amount:     10000,
			PaymentURL: "https://example.com/redirect",
			User:       userData,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:     2,
			UserID: 2,
		},
	}

	t.Run("Success get transaction by user id", func(t *testing.T) {

		repository.On("GetByUserId", 1).Return(transactions, nil).Once()

		result, err := usecase.GetTransactionByUserId(1)

		assert.NoError(t, err)

		assert.Equal(t, transactions, result)

		repository.AssertExpectations(t)
	})

	t.Run("Failed get transaction by user id", func(t *testing.T) {

		repository.On("GetByUserId", 1).Return(transactions, errors.New("error")).Once()

		result, err := usecase.GetTransactionByUserId(1)

		assert.Error(t, err)

		assert.Equal(t, transactions, result)

		repository.AssertExpectations(t)
	})

}

// func TestCreateTransaction(t *testing.T) {

// 	var repository = NewRepository(t)
// 	var orderRepository = mocks.NewRepository(t)
// 	var paymentUsecase = payment.NewUsecase()
// 	var usecase = transaction.NewUsecase(repository, orderRepository, paymentUsecase)

// 	var orderr = order.Order{
// 		ID:        uuid.New(),
// 		UserID:    1,
// 		Status:    "pending",
// 		CreatedAt: time.Now(),
// 	}

// 	var orderItems = []order.OrderItem{
// 		{
// 			OrderID:   uuid.New(),
// 			ProductID: 1,
// 			Quantity:  1,
// 		},
// 		{
// 			OrderID:   uuid.New(),
// 			ProductID: 2,
// 			Quantity:  2,
// 		},
// 	}

// 	t.Run("Success create transaction", func(t *testing.T) {

// 		orderRepository.On("GetOrderItemByID", orderr.ID).Return(orderr, nil).Once()

// 		orderRepository.On("GetOrderItemsByOrderID", orderr.ID).Return(orderItems, nil).Once()

// 		result, err := usecase.CreateTransaction(transaction.CreateTransactionInput{

// 			OrderID: orderr.ID,
// 			Amount:  10000,
// 		})

// 		assert.NoError(t, err)

// 		assert.Equal(t, transaction.Transaction{}, result)

// 		orderRepository.AssertExpectations(t)
// 	})

// 	t.Run("Failed create transaction", func(t *testing.T) {

// 		orderRepository.On("GetOrderItemByID", orderr.ID).Return(orderr, nil).Once()

// 		orderRepository.On("GetOrderItemsByOrderID", orderr.ID).Return(orderItems, errors.New("error")).Once()

// 		result, err := usecase.CreateTransaction(transaction.CreateTransactionInput{

// 			OrderID: orderr.ID,
// 			Amount:  10000,
// 		})

// 		assert.Error(t, err)

// 		assert.Equal(t, transaction.Transaction{}, result)

// 		orderRepository.AssertExpectations(t)
// 	})
// }
