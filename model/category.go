package model

import "time"

type Category struct {
	ID        int    `gorm:"primaryKey"`
	Slug      string `gorm:"size:100;not null"`
	Nama      string `gorm:"size:100;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product
}
