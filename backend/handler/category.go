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

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *CategoryHandler) GetAll(c echo.Context) error {
	categories, err := h.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch categories"})
	}
	res := make([]CategoryResponse, len(categories))
	for i, cat := range categories {
		res[i] = CategoryResponse{ID: cat.ID, Name: cat.Name}
	}
	return c.JSON(http.StatusOK, res)
}

func (h *CategoryHandler) Create(c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	category, err := h.usecase.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create category"})
	}

	return c.JSON(http.StatusOK, CategoryResponse{ID: category.ID, Name: category.Name})
}