package model

import "time"

type Role struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User
}
