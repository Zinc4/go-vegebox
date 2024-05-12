package user

type userFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) userFormatter {
	return userFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}

type VerifyEmailPayloadData struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}
