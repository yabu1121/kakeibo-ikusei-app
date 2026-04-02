package model

type Category struct {
	ID string `gorm:"primaryKey;type:uuid"`
	Name string `gorm:"not null"`
}