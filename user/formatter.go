package user

type userFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Avatar   string `json:"avatar"`
}

func FormatUser(user User, token string) userFormatter {
	return userFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Token:    token,
		Avatar:   user.Avatar,
	}
}

type VerifyEmailPayloadData struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}
