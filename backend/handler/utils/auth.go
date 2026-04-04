package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserID (c echo.Context) (string, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return "", c.JSON(http.StatusInternalServerError, map[string]string{"error": "情報の型が不正です"})
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", c.JSON(http.StatusInternalServerError, map[string]string{"error": "情報の型が不正です"})
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", c.JSON(http.StatusInternalServerError, map[string]string{"error": "情報の型が不正です"})
	}
	return userID, nil
}