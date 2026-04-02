package main

import (
	"github.com/kakebon/backend/handler"
	"github.com/kakebon/backend/infrastructure"
	"github.com/kakebon/backend/infrastructure/persistence/sqlite"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	db := infrastructure.InitDB()

	// Infrastructure層: リポジトリの実装を生成
	expenseRepo := sqlite.NewExpenseRepository(db)
	characterRepo := sqlite.NewCharacterRepository(db)

	// Usecase層: アプリケーションのビジネスロジックを生成
	expenseUsecase := usecase.NewExpenseUsecase(expenseRepo, characterRepo)
	characterUsecase := usecase.NewCharacterUsecase(characterRepo)

	// Handler層: HTTPハンドラーを生成
	expenseHandler := handler.NewExpenseHandler(expenseUsecase)
	characterHandler := handler.NewCharacterHandler(characterUsecase)

	e := echo.New()
	e.POST("/expense", expenseHandler.RecordExpense)
	e.GET("/expense", expenseHandler.GetAllExpense)
	e.GET("/character", characterHandler.GetCharacterInformation)
	e.POST("/character/login", characterHandler.LoginBonus)

	e.Start(":8080")
}
