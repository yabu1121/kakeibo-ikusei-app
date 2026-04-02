package handler

import (
	"net/http"

	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(u *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: u}
}

func (h *CategoryHandler) GetAll (c echo.Context) error {
	categories, err := h.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "fialed to fetch categories"})
	}
	return c.JSON(http.StatusOK, categories)
}