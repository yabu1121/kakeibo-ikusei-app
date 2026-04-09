package handler

import (
	"net/http"

	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/handler/utils"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

type CharacterHandler struct {
	usecase *usecase.CharacterUsecase
}

func NewCharacterHandler(u *usecase.CharacterUsecase) *CharacterHandler {
	return &CharacterHandler{usecase: u}
}

type CharacterResponse struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	CurrentLevel   int    `json:"current_level"`
	CurrentExp     int    `json:"current_exp"`
	ExpToNextLevel int    `json:"exp_to_next_level"`
	ImageURL       string `json:"image_url"`
}

func (h *CharacterHandler) GetCharacterInformation(c echo.Context) error {
	
	userID, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	char, err := h.usecase.GetByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get character"})
	}
	res := CharacterResponse{
		ID:             char.ID,
		UserID:         char.UserID,
		CurrentLevel:   char.CurrentLevel,
		CurrentExp:     char.CurrentExp,
		ExpToNextLevel: char.ExpToNextLevel,
		ImageURL:       model.GetImageByLevel(char.CurrentLevel),
	}
	return c.JSON(http.StatusOK, res)
}

func (h *CharacterHandler) LoginBonus(c echo.Context) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	char, err := h.usecase.LoginBonus(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := CharacterResponse{
		ID:             char.ID,
		UserID:         char.UserID,
		CurrentLevel:   char.CurrentLevel,
		CurrentExp:     char.CurrentExp,
		ExpToNextLevel: char.ExpToNextLevel,
		ImageURL:       model.GetImageByLevel(char.CurrentLevel),
	}
	return c.JSON(http.StatusOK, res)
}
