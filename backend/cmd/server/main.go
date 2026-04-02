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
	expensePers := sqlite.NewExpensePersistence(db)
	characterPers := sqlite.NewCharacterPersistence(db)
	categoryPers := sqlite.NewCategoryPersistence(db)
	

	// Usecase層: アプリケーションのビジネスロジックを生成
	expenseUsecase := usecase.NewExpenseUsecase(expensePers, characterPers)
	characterUsecase := usecase.NewCharacterUsecase(characterPers)
	categoryUsecase := usecase.NewCategoryUsecase(categoryPers)

	// Handler層: HTTPハンドラーを生成
	expenseHandler := handler.NewExpenseHandler(expenseUsecase)
	characterHandler := handler.NewCharacterHandler(characterUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	e := echo.New()
	e.POST("/expense", expenseHandler.RecordExpense)
	e.GET("/expense", expenseHandler.GetAllExpense)
	e.GET("/character", characterHandler.GetCharacterInformation)
	e.POST("/character/login", characterHandler.LoginBonus)
	e.GET("/category", categoryHandler.GetAll)
	e.POST("/category", categoryHandler.Create)
	
	e.Start(":8080")
}
