package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)
func CheckRole(targetRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims, ok := user.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, "invalid claims")
			}

			role, ok := claims["role"].(string)
			if !ok || role != targetRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "権限がありません"})
			}

			return next(c)
		}
	}
}