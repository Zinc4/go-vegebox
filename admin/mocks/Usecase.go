package mocks

import (
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) GetUserPagination(page, pageSize int) ([]user.User, int, int, int, int, error) {

	args := u.Called(page, pageSize)
	return args.Get(0).([]user.User), args.Get(1).(int), args.Get(2).(int), args.Get(3).(int), args.Get(4).(int), args.Error(5)
}

func (u *Usecase) SearchUserByName(name string) ([]user.User, error) {

	args := u.Called(name)
	return args.Get(0).([]user.User), args.Error(1)
}

func (u *Usecase) GetTransactionsPagination(page, pageSize int) ([]transaction.Transaction, int, int, int, int, error) {

	args := u.Called(page, pageSize)
	return args.Get(0).([]transaction.Transaction), args.Get(1).(int), args.Get(2).(int), args.Get(3).(int), args.Get(4).(int), args.Error(5)
}

func (u *Usecase) SearchTransactionByName(name string) ([]transaction.Transaction, error) {

	args := u.Called(name)
	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (u *Usecase) DeleteUserById(id int) (user.User, error) {

	args := u.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) FindUserById(id int) (user.User, error) {

	args := u.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) CreateProduct(input product.AddProductInput) (product.Product, error) {

	args := u.Called(input)
	return args.Get(0).(product.Product), args.Error(1)
}

func (u *Usecase) UpdateProduct(inputID product.GetProductDetailInput, data product.AddProductInput) (product.Product, error) {

	args := u.Called(inputID, data)
	return args.Get(0).(product.Product), args.Error(1)
}

func (u *Usecase) FindProductByID(id int) (product.Product, error) {

	args := u.Called(id)
	return args.Get(0).(product.Product), args.Error(1)
}

func (u *Usecase) FindCategoryByID(id int) (product.Category, error) {

	args := u.Called(id)
	return args.Get(0).(product.Category), args.Error(1)
}

func (u *Usecase) DeleteProductByID(id int) (product.Product, error) {

	args := u.Called(id)
	return args.Get(0).(product.Product), args.Error(1)
}

func (u *Usecase) DeleteCategoryByID(id int) (product.Category, error) {

	args := u.Called(id)
	return args.Get(0).(product.Category), args.Error(1)
}

func (u *Usecase) CreateCategory(input product.Category) (product.Category, error) {

	args := u.Called(input)
	return args.Get(0).(product.Category), args.Error(1)
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
