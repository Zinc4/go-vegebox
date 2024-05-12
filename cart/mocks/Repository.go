package mocks

import (
	"mini-project/cart"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Create(crt cart.Cart) (cart.Cart, error) {

	args := r.Called(crt)
	return args.Get(0).(cart.Cart), args.Error(1)
}

func (r *Repository) GetCartItemByCartId(cartID int) ([]cart.CartItem, error) {

	args := r.Called(cartID)
	return args.Get(0).([]cart.CartItem), args.Error(1)
}

func (r *Repository) GetById(id int) (cart.Cart, error) {

	args := r.Called(id)
	return args.Get(0).(cart.Cart), args.Error(1)
}

func (r *Repository) DeleteById(id int) error {

	args := r.Called(id)
	return args.Error(0)
}

func (r *Repository) DeleteProductByCartIdAndProductId(cartID int, productID int) error {

	args := r.Called(cartID, productID)
	return args.Error(0)
}

func (r *Repository) DeleteCartItemByCartIdAndProductId(cartID int, productID int) error {

	args := r.Called(cartID, productID)
	return args.Error(0)
}

func (r *Repository) SaveCartItem(cartItem cart.CartItem) (cart.CartItem, error) {

	args := r.Called(cartItem)
	return args.Get(0).(cart.CartItem), args.Error(1)
}

func (r *Repository) GetCartItemByCartIdAndProductId(cartID int, productID int) (cart.CartItem, error) {

	args := r.Called(cartID, productID)
	return args.Get(0).(cart.CartItem), args.Error(1)
}

func (r *Repository) GetCarts() ([]cart.Cart, error) {

	args := r.Called()
	return args.Get(0).([]cart.Cart), args.Error(1)
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
