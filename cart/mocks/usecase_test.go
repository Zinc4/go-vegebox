package mocks

import (
	"errors"
	"mini-project/cart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCart(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var cart = cart.Cart{
		ID: 1,
	}

	t.Run("success create cart", func(t *testing.T) {

		repository.On("Create", mock.AnythingOfType("cart.Cart")).Return(cart, nil).Once()

		result, err := usecase.CreateCart()

		assert.NoError(t, err)

		assert.Equal(t, cart, result)

		repository.AssertExpectations(t)
	})

	t.Run("error create cart", func(t *testing.T) {

		repository.On("Create", mock.AnythingOfType("cart.Cart")).Return(cart, errors.New("error")).Once()

		result, err := usecase.CreateCart()

		assert.Error(t, err)

		assert.Equal(t, cart, result)

		repository.AssertExpectations(t)
	})

}

func TestGetCartItemByCartId(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var cartItems = []cart.CartItem{
		{
			CartID:    1,
			ProductID: 1,
			Quantity:  1,
		},
		{
			CartID:    1,
			ProductID: 2,
			Quantity:  2,
		},
	}

	t.Run("success get cart item by cart id", func(t *testing.T) {

		repository.On("GetCartItemByCartId", mock.AnythingOfType("int")).Return(cartItems, nil).Once()

		result, err := usecase.GetCartItemByCartId(1)

		assert.NoError(t, err)

		assert.Equal(t, cartItems, result)

		repository.AssertExpectations(t)
	})

	t.Run("error get cart item by cart id", func(t *testing.T) {

		repository.On("GetCartItemByCartId", mock.AnythingOfType("int")).Return(cartItems, errors.New("error")).Once()

		result, err := usecase.GetCartItemByCartId(1)

		assert.Error(t, err)

		assert.Equal(t, cartItems, result)

		repository.AssertExpectations(t)
	})

}

func TestSaveCartItem(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var cartItem = cart.CartItem{
		CartID:    1,
		ProductID: 1,
		Quantity:  1,
	}

	t.Run("success save cart item", func(t *testing.T) {

		repository.On("SaveCartItem", mock.AnythingOfType("cart.CartItem")).Return(cartItem, nil).Once()

		result, err := usecase.SaveCartItem(cartItem)

		assert.NoError(t, err)

		assert.Equal(t, cartItem, result)

		repository.AssertExpectations(t)
	})
}

func TestGetCartItemByCartIdAndProductId(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var cartItem = cart.CartItem{
		CartID:    1,
		ProductID: 1,
		Quantity:  1,
	}

	t.Run("success get cart item by cart id and product id", func(t *testing.T) {

		repository.On("GetCartItemByCartIdAndProductId", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(cartItem, nil).Once()

		result, err := usecase.GetCartItemByCartIdAndProductId(1, 1)

		assert.NoError(t, err)

		assert.Equal(t, cartItem, result)

		repository.AssertExpectations(t)
	})

	t.Run("error get cart item by cart id and product id", func(t *testing.T) {

		repository.On("GetCartItemByCartIdAndProductId", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(cartItem, errors.New("error")).Once()

		result, err := usecase.GetCartItemByCartIdAndProductId(1, 1)

		assert.Error(t, err)

		assert.Equal(t, cartItem, result)

		repository.AssertExpectations(t)
	})
}

func TestDeleteCartItemByCartIdAndProductId(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	t.Run("success delete cart item by cart id and product id", func(t *testing.T) {

		repository.On("DeleteCartItemByCartIdAndProductId", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()

		err := usecase.DeleteCartItemByCartIdAndProductId(1, 1)

		assert.NoError(t, err)

		repository.AssertExpectations(t)
	})

	t.Run("error delete cart item by cart id and product id", func(t *testing.T) {

		repository.On("DeleteCartItemByCartIdAndProductId", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("error")).Once()

		err := usecase.DeleteCartItemByCartIdAndProductId(1, 1)

		assert.Error(t, err)

		repository.AssertExpectations(t)
	})
}

func TestDeleteCartById(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	t.Run("success delete cart by id", func(t *testing.T) {

		repository.On("DeleteById", mock.AnythingOfType("int")).Return(nil).Once()

		err := usecase.DeleteCartById(1)

		assert.NoError(t, err)

		repository.AssertExpectations(t)

	})

	t.Run("error delete cart by id", func(t *testing.T) {

		repository.On("DeleteById", mock.AnythingOfType("int")).Return(errors.New("error")).Once()

		err := usecase.DeleteCartById(1)

		assert.Error(t, err)

		repository.AssertExpectations(t)
	})
}

func TestGetCartById(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var cart = cart.Cart{
		ID: 1,
	}

	t.Run("success get cart by id", func(t *testing.T) {

		repository.On("GetById", mock.AnythingOfType("int")).Return(cart, nil).Once()

		result, err := usecase.GetCartById(1)

		assert.NoError(t, err)

		assert.Equal(t, cart, result)
	})
}

func TestGetCarts(t *testing.T) {

	repository := NewRepository(t)
	usecase := cart.NewUsecase(repository)

	var carts = []cart.Cart{
		{
			ID: 1,
		},
	}

	t.Run("success get carts", func(t *testing.T) {

		repository.On("GetCarts").Return(carts, nil).Once()

		result, err := usecase.GetCarts()

		assert.NoError(t, err)

		assert.Equal(t, carts, result)

		repository.AssertExpectations(t)
	})

	t.Run("error get carts", func(t *testing.T) {

		repository.On("GetCarts").Return([]cart.Cart{}, errors.New("error")).Once()

		result, err := usecase.GetCarts()

		assert.Error(t, err)

		assert.Equal(t, []cart.Cart{}, result)

		repository.AssertExpectations(t)
	})
}
