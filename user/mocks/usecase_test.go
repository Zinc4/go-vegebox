package mocks

import (
	"errors"
	"mini-project/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	input := user.RegisterUserInput{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password123",
	}

	t.Run("success", func(t *testing.T) {

		repository.On("Save", mock.AnythingOfType("user.User")).Return(user.User{ID: 1, Role: "user"}, nil).Once()
		repository.On("SaveOTP", mock.AnythingOfType("user.OTP")).Return(user.OTP{ID: 1}, nil).Once()

		user, err := usecase.RegisterUser(input)
		assert.NotNil(t, user.ID)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {

		repository.On("Save", mock.AnythingOfType("user.User")).Return(user.User{}, errors.New("error")).Once()

		user, err := usecase.RegisterUser(input)

		assert.Equal(t, "error", err.Error())
		assert.NotNil(t, user)
		repository.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {

	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	userData := user.User{
		ID:         1,
		Name:       "test",
		Email:      "test@example.com",
		Password:   "$argon2id$v=19$m=131072,t=4,p=4$LEJWB7alTEITjk3Z3LQa2g$mLSjwp4ThWuimkOXDYhGrwiTEfGM7pDBD8iI+CuYU8E",
		Role:       "user",
		IsVerified: true,
	}

	input := user.LoginInput{
		Email:    "test@example.com",
		Password: "password123",
	}

	t.Run("success login", func(t *testing.T) {
		repository.On("FindByEmail", input.Email).Return(userData, nil).Once()

		user, err := usecase.Login(input)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		repository.AssertExpectations(t)
	})

	t.Run("invalid email", func(t *testing.T) {

		repository.On("FindByEmail", input.Email).Return(user.User{}, errors.New("no user found with this email")).Once()

		user, err := usecase.Login(input)

		assert.NotNil(t, user)
		assert.Equal(t, "no user found with this email", err.Error())
	})
}

func TestVerifyEmail(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	email := "test@example.com"
	otp := "123456"

	t.Run("valid verification", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{ID: 1, Role: "user"}, nil).Once()
		repository.On("FindOTP", 1, otp).Return(user.OTP{ID: 1}, nil).Once()
		repository.On("UpdateUser", mock.AnythingOfType("user.User")).Return(user.User{ID: 1, IsVerified: true}, nil).Once()
		repository.On("DeleteOTP", mock.AnythingOfType("user.OTP")).Return(nil, nil).Once()

		err := usecase.VerifyEmail(email, otp)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("invalid verification", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{}, errors.New("no user found on that email")).Once()

		err := usecase.VerifyEmail(email, otp)
		assert.Equal(t, "no user found on that email", err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("invalid OTP", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{ID: 1, Role: "user"}, nil).Once()
		repository.On("FindOTP", 1, otp).Return(user.OTP{}, errors.New("invalid or expired OTP")).Once()

		err := usecase.VerifyEmail(email, otp)

		assert.Equal(t, "invalid or expired OTP", err.Error())
		repository.AssertExpectations(t)
	})
}

func TestGetUserByEmail(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	email := "test@example.com"

	t.Run("Valid get user by email", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{ID: 1, Role: "user"}, nil).Once()

		user, err := usecase.GetUserByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, user.ID)
		assert.Equal(t, user.Role, "user")
		repository.AssertExpectations(t)
	})

	t.Run("Invalid get user by email - user not found", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{}, errors.New("no user found on that email")).Once()

		_, err := usecase.GetUserByEmail(email)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "no user found on that email")
		repository.AssertExpectations(t)
	})

	t.Run("Invalid get user by email - error from repository", func(t *testing.T) {

		repository.On("FindByEmail", email).Return(user.User{}, errors.New("internal server error")).Once()

		_, err := usecase.GetUserByEmail(email)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "internal server error")
		repository.AssertExpectations(t)
	})
}

func TestFindValidOTP(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	userID := 1
	otp := "123456"

	t.Run("Valid find valid OTP", func(t *testing.T) {

		repository.On("FindOTP", userID, otp).Return(user.OTP{ID: 1}, nil).Once()

		validOTP, err := usecase.FindValidOTP(userID, otp)
		assert.NoError(t, err)
		assert.NotNil(t, validOTP.ID)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid find valid OTP - OTP not found", func(t *testing.T) {

		repository.On("FindOTP", userID, otp).Return(user.OTP{}, errors.New("invalid or expired OTP")).Once()

		_, err := usecase.FindValidOTP(userID, otp)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "invalid or expired OTP")
		repository.AssertExpectations(t)
	})

	t.Run("Invalid find valid OTP - error from repository", func(t *testing.T) {

		repository.On("FindOTP", userID, otp).Return(user.OTP{}, errors.New("internal server error")).Once()

		_, err := usecase.FindValidOTP(userID, otp)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "internal server error")
		repository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	user := user.User{
		ID:    1,
		Name:  "test",
		Email: "test@example.com",
	}

	t.Run("Valid update user", func(t *testing.T) {

		repository.On("UpdateUser", user).Return(user, nil).Once()

		updatedUser, err := usecase.UpdateUser(user)
		assert.NoError(t, err)
		assert.Equal(t, updatedUser.ID, user.ID)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid update user - user not found", func(t *testing.T) {

		repository.On("UpdateUser", user).Return(user, errors.New("user not found")).Once()

		_, err := usecase.UpdateUser(user)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "user not found")
		repository.AssertExpectations(t)
	})

	t.Run("Invalid update user - error from repository", func(t *testing.T) {

		repository.On("UpdateUser", user).Return(user, errors.New("internal server error")).Once()

		_, err := usecase.UpdateUser(user)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "internal server error")
		repository.AssertExpectations(t)
	})
}

func TestResendOTP(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	testEmail := "test@example.com"

	t.Run("Valid resend OTP", func(t *testing.T) {
		repository.On("FindByEmail", testEmail).Return(user.User{ID: 1}, nil).Once()
		repository.On("DeleteUserOTP", 1).Return(nil).Once()
		repository.On("SaveOTP", mock.AnythingOfType("user.OTP")).Return(user.OTP{ID: 1}, nil).Once()

		otp, err := usecase.ResendOTP(testEmail)

		assert.NotNil(t, otp)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid email", func(t *testing.T) {
		repository.On("FindByEmail", testEmail).Return(user.User{ID: 1}, errors.New("no user found with that email")).Once()

		_, err := usecase.ResendOTP(testEmail)

		assert.Equal(t, "no user found with that email", err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Failed to delete user OTP", func(t *testing.T) {

		repository.On("FindByEmail", testEmail).Return(user.User{ID: 1}, nil).Once()
		repository.On("DeleteUserOTP", 1).Return(nil, errors.New("failed to delete user OTP")).Once()
		repository.On("SaveOTP", mock.AnythingOfType("user.OTP")).Return(user.OTP{ID: 1}, errors.New("error")).Once()

		_, err := usecase.ResendOTP(testEmail)

		assert.EqualError(t, err, "error")
		repository.AssertExpectations(t)
	})

	t.Run("Failed to save OTP", func(t *testing.T) {
		repository.On("FindByEmail", testEmail).Return(user.User{ID: 1}, nil).Once()
		repository.On("DeleteUserOTP", 1).Return(nil).Once()
		repository.On("SaveOTP", mock.AnythingOfType("user.OTP")).Return(user.OTP{ID: 1}, errors.New("failed to save OTP")).Once()

		_, err := usecase.ResendOTP(testEmail)

		assert.EqualError(t, err, "failed to save OTP")
		repository.AssertExpectations(t)
	})
}

// func TestSaveAvatar(t *testing.T) {
// 	repository := NewRepository(t)
// 	usecase := user.NewUsecase(repository)

// 	userID := 1
// 	file := "avatar.png"

// 	t.Run("Valid save avatar", func(t *testing.T) {

// 		users := user.User{ID: userID}
// 		repository.On("FindByID", userID).Return(users, nil).Once()
// 		updatedUser := user.User{ID: userID, Avatar: file}
// 		repository.On("Update", updatedUser).Return(updatedUser, nil).Once()

// 		updatedUser, err := usecase.SaveAvatar(userID, file)
// 		assert.NoError(t, err)
// 		assert.Equal(t, updatedUser.ID, userID)
// 		assert.Equal(t, updatedUser.Avatar, file)
// 		repository.AssertExpectations(t)
// 	})

// 	t.Run("Invalid save avatar - user not found", func(t *testing.T) {

// 		repository.On("FindByID", userID).Return(user.User{}, errors.New("no user found with that ID")).Once()

// 		updatedUser, err := usecase.SaveAvatar(userID, file)

// 		assert.Error(t, err)
// 		assert.Equal(t, err.Error(), "no user found with that ID")
// 		assert.Equal(t, user.User{}, updatedUser)
// 		repository.AssertExpectations(t)
// 	})

// 	t.Run("Invalid save avatar - error updating user", func(t *testing.T) {

// 		users := user.User{ID: userID, Avatar: file}
// 		repository.On("FindByID", userID).Return(users, nil).Once()
// 		repository.On("Update", users).Return(users, errors.New("error updating user")).Once()

// 		updatedUser, err := usecase.SaveAvatar(userID, file)

// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "error updating user")
// 		assert.NotEqual(t, user.User{}, updatedUser)
// 		repository.AssertExpectations(t)
// 	})
// }

func TestGetUserByID(t *testing.T) {
	repository := NewRepository(t)
	usecase := user.NewUsecase(repository)

	userID := 1

	t.Run("Valid get user by ID", func(t *testing.T) {

		user := user.User{ID: userID}
		repository.On("FindByID", userID).Return(user, nil).Once()

		actualUser, err := usecase.GetUserByID(userID)
		assert.NoError(t, err)
		assert.Equal(t, user, actualUser)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid get user by ID - user not found", func(t *testing.T) {

		repository.On("FindByID", userID).Return(user.User{}, errors.New("no user found on with that ID")).Once()

		actualUser, err := usecase.GetUserByID(userID)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "no user found on with that ID")
		assert.Equal(t, user.User{}, actualUser)
		repository.AssertExpectations(t)
	})
}
