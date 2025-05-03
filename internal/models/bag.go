package models

import "gorm.io/gorm"

type Bag struct {
	gorm.Model
	Title    string  `json:"title" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Color    string  `json:"color" binding:"required"`
	Price    float64 `json:"price" binding:"required,gte=0"`
	Size     string  `json:"size" binding:"required"`

	UserID uint `json:"user_id"`                                                // внешний ключ
	User   User `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // связь с пользователем
}
