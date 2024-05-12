package mocks

import (
	"mini-project/transaction"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetByOrderId(id uuid.UUID) ([]transaction.Transaction, error) {

	args := r.Called(id)

	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (r *Repository) GetByUserId(id int) ([]transaction.Transaction, error) {

	args := r.Called(id)

	return args.Get(0).([]transaction.Transaction), args.Error(1)
}

func (r *Repository) Save(transac transaction.Transaction) (transaction.Transaction, error) {

	args := r.Called(transac)

	return args.Get(0).(transaction.Transaction), args.Error(1)
}

func (r *Repository) Update(transac transaction.Transaction) (transaction.Transaction, error) {

	args := r.Called(transac)

	return args.Get(0).(transaction.Transaction), args.Error(1)
}

func (r *Repository) GetByCode(code string) (transaction.Transaction, error) {

	args := r.Called(code)

	return args.Get(0).(transaction.Transaction), args.Error(1)
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
