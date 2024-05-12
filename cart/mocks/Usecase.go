package mocks

import (
	"mini-project/cart"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) CreateCart() (cart.Cart, error) {

	args := u.Called()
	return args.Get(0).(cart.Cart), args.Error(1)
}

func (u *Usecase) GetCartItemByCartId(cartID int) ([]cart.CartItem, error) {

	args := u.Called(cartID)
	return args.Get(0).([]cart.CartItem), args.Error(1)
}

func (u *Usecase) SaveCartItem(cartItem cart.CartItem) (cart.CartItem, error) {

	args := u.Called(cartItem)
	return args.Get(0).(cart.CartItem), args.Error(1)
}

func (u *Usecase) GetCartItemByCartIdAndProductId(cartID int, productID int) (cart.CartItem, error) {

	args := u.Called(cartID, productID)
	return args.Get(0).(cart.CartItem), args.Error(1)
}

func (u *Usecase) DeleteCartItemByCartIdAndProductId(cartID int, productID int) error {

	args := u.Called(cartID, productID)
	return args.Error(0)
}

func (u *Usecase) DeleteCartById(cartID int) error {

	args := u.Called(cartID)
	return args.Error(0)
}

func (u *Usecase) GetCartById(id int) (cart.Cart, error) {

	args := u.Called(id)
	return args.Get(0).(cart.Cart), args.Error(1)
}

func (u *Usecase) GetCarts() ([]cart.Cart, error) {

	args := u.Called()
	return args.Get(0).([]cart.Cart), args.Error(1)
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
