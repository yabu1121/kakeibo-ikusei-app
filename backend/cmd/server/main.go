package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func hello (c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

type Expense struct {
	ID uint `gorm:"primaryKey"`
	Amount int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ExpenseResponse struct {
	Amount int `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCharacter struct {
	ID uint `gorm:"primaryKey"`
	CurrentLevel int
	CurrentExp int
	ExpToNextLevel int
	ImageURL string
}

type CharacterResponse struct {
	CurrentLevel int `json:"current_level"`
	CurrentExp int `json:"current_exp"`
	ExpToNextLevel int `json:"exp_to_next_level"`
	ImageURL string `json:"image_url"`
}

type ExpenseRequest struct {
	Amount int `json:"amount"`
}

type ExpenseHandler struct {
	DB *gorm.DB
}

type CharacterHandler struct {
	DB *gorm.DB
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect sqlite db", err)
	}
	db.AutoMigrate(&Expense{}, &UserCharacter{})
	var count int64
	db.Model(&UserCharacter{}).Count(&count)
	if count == 0 {
		db.Create(&UserCharacter{
			CurrentLevel:   1,
			CurrentExp:     0,
			ExpToNextLevel: 100, 
			ImageURL:       "https://example.com/images/level1_egg.png",
		})
		log.Println("初期キャラクター（Lv1）のデータをDBに作成しました")
	}

	return db
}

func CalcExp (char *UserCharacter, exp int) {
	char.CurrentExp += exp
	for char.CurrentExp >= char.ExpToNextLevel {
		char.CurrentExp -= char.ExpToNextLevel
		char.CurrentLevel++
		char.ExpToNextLevel = char.CurrentLevel * 100
		if char.CurrentLevel >= 99 {
			char.CurrentLevel = 99
			char.ExpToNextLevel = 9900
		}
		if char.CurrentLevel >= 2 {
			char.ImageURL = "https://example.com/images/level2_bird.png"
		}
	}
}

func (h *ExpenseHandler)RecordExpense(c echo.Context) error {
	var req ExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	
	var char UserCharacter
	err :=  h.DB.Transaction(func (tx *gorm.DB) error {
		expense := Expense{Amount: req.Amount}
		if err := tx.Create(&expense).Error; err != nil {
			return err 
		}
		
		if err := tx.First(&char).Error; err != nil {
			return err 
		}

		getExp := req.Amount / 100 
		CalcExp(&char, getExp)

		if err := tx.Save(&char).Error; err != nil {
			return err 
		}
		return nil
	})

	if err != nil {
		log.Printf("Failed to record expense: %v", err) 
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, char)
}

func (h *ExpenseHandler)GetAllExpense(c echo.Context) error {
	var expense []Expense
	if err := h.DB.Find(&expense).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get expense list"})
	}
	return c.JSON(http.StatusOK, expense)
}

func (h *CharacterHandler) GetCharacterInfomation (c echo.Context) error {
	var uc UserCharacter
	if err := h.DB.First(&uc).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get character"})
	}
	res := CharacterResponse{
		CurrentLevel: uc.CurrentLevel,
		CurrentExp: uc.CurrentExp,
		ExpToNextLevel: uc.ExpToNextLevel,
		ImageURL: uc.ImageURL,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *CharacterHandler) LoginBonus (c echo.Context) error {
	getExp := 5
	var char UserCharacter
	if err := h.DB.First(&char).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get user character"})
	}

	CalcExp(&char, getExp)
	if err := h.DB.Save(&char).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, char)
}
// func (h *ExpenseHandler)GetExpenseById(c echo.Context) error {}
// func (h *ExpenseHandler)DeleteExpense(c echo.Context) error {}
// func (h *ExpenseHandler)UpdateExpense(c echo.Context) error {}

func main() {
	e := echo.New();

	db := InitDB()
	handler := &ExpenseHandler{ DB: db }
	chanlder := &CharacterHandler{ DB: db }

	e.GET("/", hello)

	e.POST("/expense", handler.RecordExpense)
	e.GET("/expense", handler.GetAllExpense)
	// e.GET("/expense/:id", hello)
	// e.PUT("/expense/:id", hello)
	// e.DELETE("/expense/:id", hello)
	e.GET("/character", chanlder.GetCharacterInfomation)
	e.POST("/character/login", chanlder.LoginBonus)
	e.Start((":8080"))
}
