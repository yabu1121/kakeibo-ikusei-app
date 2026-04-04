package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler (u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

type createRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var req createRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.usecase.Create(req.Name, req.Email, req.Password);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email query param is required"})
	}

	user, err := h.usecase.GetByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET_KEY")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}