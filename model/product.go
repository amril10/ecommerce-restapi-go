package model

import (
	"time"
)

type Product struct {
	ID          int    `gorm:"primaryKey"`
	Slug        string `gorm:"size:150;not null"`
	CategoryID  int
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
	Name        string   `gorm:"size:150;not null"`
	Description string
	Price       int64 `gorm:"not null"`
	CoverUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Carts       []Cart
}
