package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserID (c echo.Context) (string, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return "", errors.New("error")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("error")
	}
	return userID, nil
}