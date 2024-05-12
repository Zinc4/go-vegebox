package middleware

import (
	"mini-project/auth"
	"mini-project/helper"
	"mini-project/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authUsecase auth.Usecase, userUsecase user.Usecase, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		authHeader := req.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.GeneralResponse("Unauthorized 1")
			return c.JSON(http.StatusUnauthorized, response)
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authUsecase.ValidateToken(tokenString)
		if err != nil {
			response := helper.GeneralResponse("Unauthorized 2")
			return c.JSON(http.StatusUnauthorized, response)

		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.GeneralResponse("Unauthorized 3")
			return c.JSON(http.StatusUnauthorized, response)

		}

		userID := int(claim["user_id"].(float64))

		user, err := userUsecase.GetUserByID(userID)
		if err != nil {
			response := helper.GeneralResponse("Unauthorized 5")
			return c.JSON(http.StatusUnauthorized, response)

		}
		c.Set("CurrentUser", user)

		return next(c)

	}

}
