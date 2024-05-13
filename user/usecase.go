package user

import (
	"errors"
	"mini-project/helper"
	"time"

	"github.com/alexedwards/argon2id"
)

type Usecase interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByEmail(email string) (User, error)
	FindValidOTP(userID int, otp string) (OTP, error)
	UpdateUser(user User) (User, error)
	VerifyEmail(email string, otp string) error
	ResendOTP(email string) (OTP, error)
	SaveAvatar(userID int, file string) (User, error)
	GetUserByID(ID int) (User, error)
	UpdateName(input UpdateProfile) (User, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	if input.Name == "" {
		return user, errors.New("name cannot be empty")
	}
	if input.Email == "" {
		return user, errors.New("email cannot be empty")
	}
	if input.Password == "" {
		return user, errors.New("password cannot be empty")
	}

	passwordHash, err := argon2id.CreateHash(input.Password, &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		return user, err
	}
	user.Password = passwordHash
	user.Role = "user"
	user.Avatar = "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1715517059/default-avatar_yt6eua.png"

	if user.Name == "" && user.Email == "" && user.Password == "" {
		return user, errors.New("please fill your name, email, and password")
	}

	newUser, err := u.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	otp := helper.GenerateRandomOTP(6)

	otpModel := OTP{
		UserId:     newUser.ID,
		ExpiredOTP: time.Now().Add(time.Minute * 10).Unix(),
		OTP:        otp,
	}

	_, errOtp := u.repository.SaveOTP(otpModel)
	if errOtp != nil {
		return newUser, errOtp
	}

	err = helper.SendOTPByEmail(newUser.Email, otp)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (u *usecase) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	if !user.IsVerified {
		return user, errors.New("user not verified")
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return user, err
	}

	if !match {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (u *usecase) GetUserByEmail(email string) (User, error) {
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *usecase) FindValidOTP(userID int, otp string) (OTP, error) {

	otpData, err := u.repository.FindOTP(userID, otp)
	if err != nil {
		return otpData, err
	}
	return otpData, nil
}

func (u *usecase) UpdateUser(user User) (User, error) {

	updatedUser, err := u.repository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *usecase) VerifyEmail(email string, otp string) error {
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("user not found on that email")
	}

	otpModel, err := u.repository.FindOTP(user.ID, otp)
	if err != nil {
		return errors.New("invalid or expired OTP")
	}

	if otpModel.ID == 0 {
		return errors.New("invalid or expired OTP")
	}

	user.IsVerified = true

	_, errUpdate := u.repository.UpdateUser(user)
	if errUpdate != nil {
		return errors.New("failed to update user")
	}

	errDeleteOTP := u.repository.DeleteOTP(otpModel)
	if errDeleteOTP != nil {
		return errors.New("failed to delete OTP")
	}

	return nil

}

func (u *usecase) ResendOTP(email string) (OTP, error) {

	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return OTP{}, err
	}

	if user.ID == 0 {
		return OTP{}, errors.New("user not found on that email")
	}

	errDeleteOTP := u.repository.DeleteUserOTP(user.ID)
	if errDeleteOTP != nil {
		return OTP{}, errDeleteOTP
	}

	otp := helper.GenerateRandomOTP(6)
	otpModel := OTP{
		UserId:     user.ID,
		ExpiredOTP: time.Now().Add(time.Minute * 10).Unix(),
		OTP:        otp,
	}

	_, errOtp := u.repository.SaveOTP(otpModel)
	if errOtp != nil {
		return OTP{}, errOtp
	}

	return otpModel, nil
}

func (u *usecase) GetUserByID(ID int) (User, error) {

	user, err := u.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (s *usecase) SaveAvatar(userID int, file string) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	user.Avatar = file

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *usecase) UpdateName(input UpdateProfile) (User, error) {

	user, err := u.repository.FindByID(input.ID)

	if user.Name == input.Name {
		return user, errors.New("name already exists")
	}

	user.Name = input.Name

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	updatedUser, err := u.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}
