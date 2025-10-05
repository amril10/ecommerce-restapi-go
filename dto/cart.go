package dto

import "time"

type CartRequest struct {
	ProductID int  `json:"product_id" binding:"required"`
	Qty       uint `json:"qty" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	Qty      uint `json:"qty" binding:"required,min=1"`
	IsSelect bool `json:"is_select"`
}

type CartResponse struct {
	ID        int  `json:"id"`
	ProductID int  `json:"product_id"`
	UserID    int  `json:"user_id"`
	Qty       uint `json:"qty"`
	Total     int  `json:"total"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateCartResponse struct {
	ID        int  `json:"id"`
	ProductID int  `json:"product_id"`
	UserID    int  `json:"user_id"`
	Qty       uint `json:"qty"`
	Total     int  `json:"total"`
	IsSelect  bool `json:"is_select"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
