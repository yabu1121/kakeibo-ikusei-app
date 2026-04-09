package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kakebon/backend/handler"
	"github.com/kakebon/backend/handler/utils"
	"github.com/kakebon/backend/infrastructure"
	"github.com/kakebon/backend/infrastructure/persistence/sqlite"
	"github.com/kakebon/backend/infrastructure/slack"
	"github.com/kakebon/backend/usecase"
	"github.com/labstack/echo-jwt/v4"
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
	userRepo := sqlite.NewUserPersistence(db)

	// Usecase層: アプリケーションのビジネスロジックを生成
	expenseUsecase := usecase.NewExpenseUsecase(expensePers, characterPers, slackRepo)
	characterUsecase := usecase.NewCharacterUsecase(characterPers)
	categoryUsecase := usecase.NewCategoryUsecase(categoryPers)
	slackUsecase := usecase.NewSlackUsecase(slackRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Handler層: HTTPハンドラーを生成
	expenseHandler := handler.NewExpenseHandler(expenseUsecase)
	characterHandler := handler.NewCharacterHandler(characterUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	slackHandler := handler.NewSlackHandler(slackUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	e.POST("/login", userHandler.Login)
	e.POST("/signup", userHandler.SignUp)


	secretKey := os.Getenv("JWT_SECRET_KEY")
	user := e.Group("/user")
	user.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	}))
	user.Use(utils.CheckRole("user"))
	
	user.POST("/expense", expenseHandler.RecordExpense)
	user.DELETE("/expense/:id", expenseHandler.DeleteByID)
	user.GET("/expense/:id", expenseHandler.GetByID)
	user.PUT("/expense/:id", expenseHandler.Update)

	user.GET("/character", characterHandler.GetCharacterInformation)
	user.POST("/character/login", characterHandler.LoginBonus)

	user.GET("/category", categoryHandler.GetAll)
	user.POST("/category", categoryHandler.Create)

	user.POST("/slack/notify", slackHandler.Notify)
	

	admin := e.Group("/admin")
	admin.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	}))
	admin.Use(utils.CheckRole("admin"))
	admin.GET("/expense", expenseHandler.GetAllExpense)


	e.Start(":8080")
}
