package mocks

import (
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetTotalUsers() (int64, error) {

	args := r.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (r *Repository) GetPaginatedUsers(offset int, limit int) ([]user.User, error) {

	args := r.Called(offset, limit)
	return args.Get(0).([]user.User), args.Error(1)
}

func (r *Repository) SearchUserByName(name string) ([]user.User, error) {

	args := r.Called(name)
	return args.Get(0).([]user.User), args.Error(1)
}

func (r *Repository) GetUserByID(userID int) (user.User, error) {

	args := r.Called(userID)
	return args.Get(0).(user.User), args.Error(1)
}

func (r *Repository) FindUserById(userID int) (user.User, error) {

	args := r.Called(userID)
	return args.Get(0).(user.User), args.Error(1)
}

func (r *Repository) GetTotalTransaction() (int64, error) {

	args := r.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (r *Repository) GetPaginatedTransaction(offset int, limit int) ([]transaction.Transaction, error) {

	args := r.Called(offset, limit)
	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (r *Repository) SearchTransactionByName(name string) ([]transaction.Transaction, error) {

	args := r.Called(name)
	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (r *Repository) GetTotalTransactionByName(name string) (int64, error) {

	args := r.Called(name)
	return args.Get(0).(int64), args.Error(1)
}

func (r *Repository) Save(prdt product.Product) (product.Product, error) {
	args := r.Called(prdt)
	return args.Get(0).(product.Product), args.Error(1)
}

func (r *Repository) Update(prdt product.Product) (product.Product, error) {

	args := r.Called(prdt)
	return args.Get(0).(product.Product), args.Error(1)
}

func (r *Repository) FindProductByID(ID int) (product.Product, error) {

	args := r.Called(ID)
	return args.Get(0).(product.Product), args.Error(1)
}

func (r *Repository) FindByCategoryByID(ID int) (product.Category, error) {

	args := r.Called(ID)
	return args.Get(0).(product.Category), args.Error(1)
}

func (r *Repository) GetProductByID(ID int) (product.Product, error) {

	args := r.Called(ID)
	return args.Get(0).(product.Product), args.Error(1)
}

func (r *Repository) GetCategoryByID(ID int) (product.Category, error) {

	args := r.Called(ID)
	return args.Get(0).(product.Category), args.Error(1)
}

func (r *Repository) SaveCategory(category product.Category) (product.Category, error) {

	args := r.Called(category)
	return args.Get(0).(product.Category), args.Error(1)
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
