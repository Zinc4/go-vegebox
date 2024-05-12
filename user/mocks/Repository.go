package mocks

import (
	"mini-project/user"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Save(u user.User) (user.User, error) {
	args := m.Called(u)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *RepositoryMock) FindByEmail(email string) (user.User, error) {
	args := m.Called(email)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *RepositoryMock) SaveOTP(otp user.OTP) (user.OTP, error) {
	args := m.Called(otp)
	return args.Get(0).(user.OTP), args.Error(1)
}

func (m *RepositoryMock) FindOTP(userID int, otp string) (user.OTP, error) {
	args := m.Called(userID, otp)
	return args.Get(0).(user.OTP), args.Error(1)
}

func (m *RepositoryMock) UpdateUser(us user.User) (user.User, error) {
	args := m.Called(us)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *RepositoryMock) DeleteOTP(otp user.OTP) error {
	args := m.Called(otp)
	return args.Error(0)
}

func (m *RepositoryMock) DeleteUserOTP(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *RepositoryMock) FindByID(id int) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *RepositoryMock) ExistingUser(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *RepositoryMock) Update(us user.User) (user.User, error) {
	args := m.Called(us)
	return args.Get(0).(user.User), args.Error(1)
}

func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryMock {
	mock := &RepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
