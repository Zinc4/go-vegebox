package user

type RegisterUserInput struct {
	Name     string `json:"name,omitempty" form:"name,omitempty" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" form:"email,omitempty" binding:"required"`
	Role     string `json:"role,omitempty" `
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ResendOTPInput struct {
	Email string `json:"email" form:"email" binding:"required"`
}
