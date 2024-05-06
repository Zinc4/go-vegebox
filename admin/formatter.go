package admin

import (
	"mini-project/user"
	"time"
)

type UserFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatUser(user user.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Avatar = user.Avatar
	userFormatter.Role = user.Role
	userFormatter.IsVerified = user.IsVerified
	userFormatter.CreatedAt = user.CreateAt

	return userFormatter
}

func FormatterUsers(users []user.User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		formatUser := FormatUser(user)
		usersFormatter = append(usersFormatter, formatUser)
	}

	return usersFormatter
}
