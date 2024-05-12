package mocks

import (
	"mini-project/user"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (u *Usecase) RegisterUser(input user.RegisterUserInput) (user.User, error) {
	args := u.Called(input)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) Login(input user.LoginInput) (user.User, error) {
	args := u.Called(input)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) GetUserByEmail(email string) (user.User, error) {
	args := u.Called(email)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) FindValidOTP(userID int, otp string) (user.OTP, error) {
	args := u.Called(userID, otp)
	return args.Get(0).(user.OTP), args.Error(1)
}

func (u *Usecase) UpdateUser(usr user.User) (user.User, error) {
	args := u.Called(usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) VerifyEmail(email string, otp string) error {
	args := u.Called(email, otp)
	return args.Error(0)
}

func (u *Usecase) ResendOTP(email string) (user.OTP, error) {
	args := u.Called(email)
	return args.Get(0).(user.OTP), args.Error(1)
}

func (u *Usecase) SaveAvatar(userID int, file string) (user.User, error) {
	args := u.Called(userID, file)
	return args.Get(0).(user.User), args.Error(1)
}

func (u *Usecase) GetUserByID(ID int) (user.User, error) {
	args := u.Called(ID)
	return args.Get(0).(user.User), args.Error(1)
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
