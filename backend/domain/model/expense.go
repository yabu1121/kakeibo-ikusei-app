package model

import "time"

type Expense struct {
	ID         string    `gorm:"primaryKey;type:uuid"`
	Name       string    `gorm:"not null"`
	Amount     int       `gorm:"not null"`
	OccuredAt  time.Time `gorm:"index;not null"`
	CategoryID string    `gorm:"index;not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	UserID     string    `gorm:"index;not null"`
	Timestamps
}