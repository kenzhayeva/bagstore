package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null;default:''"`
	Role      string    `gorm:"not null;default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Bags []Bag `gorm:"foreignKey:UserID"` // Один пользователь — много сумок
}
