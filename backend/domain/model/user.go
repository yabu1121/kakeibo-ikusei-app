package model

type User struct {
	ID             string     `gorm:"primaryKey;type:uuid"`
	Name           string     `gorm:"not null"`
	Email          string     `gorm:"uniqueIndex;not null"`
	HashedPassword string     `gorm:"not null"`
	Role           string     `gorm:"not null"`
	Character      *Character `gorm:"foreignKey:UserID"`
	Timestamps
}