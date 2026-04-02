package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kakebon/backend/handler"
	"github.com/kakebon/backend/infrastructure"
	"github.com/kakebon/backend/infrastructure/persistence/sqlite"
	"github.com/kakebon/backend/infrastructure/slack"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	// go run ./cmd/server/ (from backend/) と go run . (from cmd/server/) の両方に対応
	for _, path := range []string{".env", "../../.env"} {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}

	db := infrastructure.InitDB()
	slackURL := os.Getenv("SLACK_WEBHOOK_URL")

	// Infrastructure層: リポジトリの実装を生成
	expensePers := sqlite.NewExpensePersistence(db)
	characterPers := sqlite.NewCharacterPersistence(db)
	categoryPers := sqlite.NewCategoryPersistence(db)
	slackRepo := slack.NewSlackNotifier(slackURL)

	// Usecase層: アプリケーションのビジネスロジックを生成
	expenseUsecase := usecase.NewExpenseUsecase(expensePers, characterPers, slackRepo)
	characterUsecase := usecase.NewCharacterUsecase(characterPers)
	categoryUsecase := usecase.NewCategoryUsecase(categoryPers)
	slackUsecase := usecase.NewSlackUsecase(slackRepo)

	// Handler層: HTTPハンドラーを生成
	expenseHandler := handler.NewExpenseHandler(expenseUsecase)
	characterHandler := handler.NewCharacterHandler(characterUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	slackHandler := handler.NewSlackHandler(slackUsecase)

	e := echo.New()
	e.POST("/expense", expenseHandler.RecordExpense)
	e.GET("/expense", expenseHandler.GetAllExpense)
	e.DELETE("/expense/:id", expenseHandler.DeleteByID)
	e.GET("/expense/:id", expenseHandler.GetByID)
	e.GET("/character", characterHandler.GetCharacterInformation)
	e.POST("/character/login", characterHandler.LoginBonus)
	e.GET("/category", categoryHandler.GetAll)
	e.POST("/category", categoryHandler.Create)
	e.POST("/slack/notify", slackHandler.Notify)

	e.Start(":8080")
}
