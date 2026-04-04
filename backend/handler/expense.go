package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/handler/utils"
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
	Name       string    `json:"name"`
	Amount     int       `json:"amount"`
	CategoryID string    `json:"category_id"`
	OccuredAt  time.Time `json:"occured_at"`
}

func (h *ExpenseHandler) RecordExpense(c echo.Context) error {
	var req ExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "amount must be positive"})
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	expense := &model.Expense{
		ID:         uuid.New().String(),
		Name:       req.Name,
		Amount:     req.Amount,
		OccuredAt: 	req.OccuredAt,
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

func (h *ExpenseHandler) DeleteByID (c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.usecase.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete expense"})
	}
	return c.NoContent(http.StatusNoContent)
} 

func (h *ExpenseHandler) GetByID (c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	res, err := h.usecase.GetByID(id); 
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "failed to get expense by this id"})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ExpenseHandler) Update (c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	var req ExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	new_expense := &model.Expense{
		ID: id,
		Name: req.Name,
		Amount: req.Amount,
		OccuredAt: req.OccuredAt,
		CategoryID: req.CategoryID,
		UserID: userID,
	}
	
	res, err := h.usecase.Update(id, new_expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update expense"})
	}
	return c.JSON(http.StatusOK, res)
}