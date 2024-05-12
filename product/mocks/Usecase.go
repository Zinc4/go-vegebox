package mocks

import (
	"mini-project/product"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) FindProducts() ([]product.Product, error) {

	args := u.Called()

	return args.Get(0).([]product.Product), args.Error(1)
}

func (u *Usecase) FindProductByID(ID int) (product.Product, error) {

	args := u.Called()

	return args.Get(0).(product.Product), args.Error(1)
}

func (u *Usecase) FindAllCategory() ([]product.Category, error) {

	args := u.Called()

	return args.Get(0).([]product.Category), args.Error(1)
}

func (u *Usecase) FindProductByCategory(categoryID int) ([]product.Product, error) {

	args := u.Called()

	return args.Get(0).([]product.Product), args.Error(1)
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
