package model

type Character struct {
	ID             string `gorm:"primaryKey;type:uuid"`
	UserID         string `gorm:"uniqueIndex;not null"`
	CurrentLevel   int    `gorm:"default:1"`
	CurrentExp     int    `gorm:"default:0"`
	ExpToNextLevel int    `gorm:"default:100"`
	ImageURL       string
}
