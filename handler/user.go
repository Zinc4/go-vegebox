package handler

import (
	"context"
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/user"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase user.Usecase
	authUsecase auth.Usecase
}

func NewUserHandler(userUsecase user.Usecase, authUsecase auth.Usecase) *userHandler {
	return &userHandler{userUsecase, authUsecase}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	input := user.RegisterUserInput{}
	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		response := helper.ErrorResponse("failed to register account", errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	// _, err = h.userUsecase.GetUserByEmail(input.Email)
	// if err != nil {
	// 	response := helper.ErrorResponse("failed to register account", err.Error())
	// 	return c.JSON(http.StatusConflict, response)
	// }

	_, err = h.userUsecase.RegisterUser(input)
	if err != nil {
		response := helper.ErrorResponse("failed to register account", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("success register account, please check your email for verification"))
}

func (h *userHandler) LoginUser(c echo.Context) error {
	var input user.LoginInput
	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		response := helper.ErrorResponse("failed to login", errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	loggedUser, err := h.userUsecase.Login(input)
	if err != nil {
		response := helper.ErrorResponse("failed to login", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authUsecase.GenerateToken(loggedUser.ID)
	if err != nil {
		response := helper.ErrorResponse("failed to login", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := user.FormatUser(loggedUser, token)

	response := helper.ResponseWithData("successfully login", formatter)

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) VerifyEmail(c echo.Context) error {
	var input user.VerifyEmailPayloadData

	err := c.Bind(&input)
	if err != nil {
		response := helper.ErrorResponse("failed to verify email", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	err = h.userUsecase.VerifyEmail(input.Email, input.OTP)
	if err != nil {
		response := helper.ErrorResponse("failed to verify email", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.SuccesResponse("success verify email")

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) ResendOTP(c echo.Context) error {
	var input user.ResendOTPInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("failed to resend OTP", errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	otp, err := h.userUsecase.ResendOTP(input.Email)
	if err != nil {
		response := helper.ErrorResponse("failed to resend OTP", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = helper.SendOTPByEmail(input.Email, otp.OTP)
	if err != nil {
		response := helper.ErrorResponse("failed to resend OTP", err.Error())
		return c.JSON(http.StatusInternalServerError, response)

	}

	response := helper.SuccesResponse("success resend OTP")
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)

	if err := c.Bind(&currentUser.ID); err != nil {
		response := helper.ErrorResponse("Error payload", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	fileHeader, _ := c.FormFile("avatar")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cloudinaryUsecase, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cloudinaryUsecase.Upload.Upload(ctx, file, uploader.UploadParams{})

	_, err := h.userUsecase.SaveAvatar(currentUser.ID, response.SecureURL)
	if err != nil {
		response := helper.ErrorResponse("failed to upload avatar", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	res := helper.UpdateAvatarRes("success upload avatar", response.SecureURL)

	return c.JSON(http.StatusOK, res)

}

func (h *userHandler) UpdateProfile(c echo.Context) error {
	currentUser := c.Get("CurrentUser").(user.User)

	var input user.UpdateProfile

	input.ID = currentUser.ID

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("failed to update profile", errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	updatedUser, err := h.userUsecase.UpdateName(input)
	if err != nil {
		response := helper.ErrorResponse("failed to update profile", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ResponseWithData("success update profile", user.FormatUserProfile(updatedUser))

	return c.JSON(http.StatusOK, response)

}
