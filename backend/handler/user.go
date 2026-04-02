package handler

import (
	"net/http"

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