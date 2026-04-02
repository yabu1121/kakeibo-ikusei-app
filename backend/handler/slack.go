package handler

import (
	"net/http"

	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

type SlackHandler struct {
	usecase *usecase.SlackUsecase
}

func NewSlackHandler (u *usecase.SlackUsecase) *SlackHandler {
	return &SlackHandler{usecase: u}
}

func (h *SlackHandler) Notify (c echo.Context) error {
	var req struct {
		Message string `json:"message"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.usecase.Execute(req.Message); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to send message to slack"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}