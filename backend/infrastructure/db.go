package infrastructure

import (
	"log"

	"github.com/kakebon/backend/domain/model"
	gormSqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(gormSqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect sqlite db", err)
	}
	db.AutoMigrate(
		&model.User{}, 
		&model.Character{}, 
		&model.Category{}, 
		&model.Expense{})
	return db
}