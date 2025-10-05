package model

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	RoleID    int
	Role      Role   `gorm:"foreignKey:RoleID;references:ID"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;size:100;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Carts     []Cart
}
