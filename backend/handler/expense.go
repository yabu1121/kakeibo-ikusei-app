package handler

import (
	"net/http"

	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct {
	usecase *usecase.ExpenseUsecase
}

func NewExpenseHandler(u *usecase.ExpenseUsecase) *ExpenseHandler {
	return &ExpenseHandler{usecase: u}
}

type ExpenseRequest struct {
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	CategoryID string `json:"category_id"`
}

func (h *ExpenseHandler) RecordExpense(c echo.Context) error {
	var req ExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "amount must be positive"})
	}

	// TODO: JWT認証後はトークンからuserIDを取得する
	userID := "dummy-user-id"

	expense := &model.Expense{
		Name:       req.Name,
		Amount:     req.Amount,
		CategoryID: req.CategoryID,
		UserID:     userID,
	}

	char, err := h.usecase.RecordExpense(expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return c.JSON(http.StatusOK, char)
}

func (h *ExpenseHandler) GetAllExpense(c echo.Context) error {
	expenses, err := h.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get expenses"})
	}
	return c.JSON(http.StatusOK, expenses)
}
