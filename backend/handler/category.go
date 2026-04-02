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

type CreateRequest struct {
	Name string `json:"name"`
}

func (h *CategoryHandler) GetAll (c echo.Context) error {
	categories, err := h.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch categories"})
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Create (c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	category, err := h.usecase.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create category"})
	}
	
	return c.JSON(http.StatusOK, category)
}