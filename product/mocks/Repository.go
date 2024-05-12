package mocks

import (
	"mini-project/product"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) FindAll() ([]product.Product, error) {

	args := r.Called()

	return args.Get(0).([]product.Product), args.Error(1)

}

func (r *Repository) FindByID(ID int) (product.Product, error) {

	args := r.Called(ID)

	return args.Get(0).(product.Product), args.Error(1)
}

func (r *Repository) FindAllCategory() ([]product.Category, error) {

	args := r.Called()

	return args.Get(0).([]product.Category), args.Error(1)
}

func (r *Repository) FindByCategory(id int) ([]product.Product, error) {

	args := r.Called(id)

	return args.Get(0).([]product.Product), args.Error(1)
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
