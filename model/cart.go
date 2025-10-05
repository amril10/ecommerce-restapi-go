package model

import "time"

type Cart struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	UserID    int
	Qty       uint
	Total     int
	IsSelect  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
	User      User    `gorm:"foreignKey:UserID;references:ID"`
}
