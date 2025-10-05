package dto

import "time"

type ProductRequest struct {
	Slug        string `json:"slug" binding:"required,min=3,max=150"`
	CategoryID  int    `json:"category_id" binding:"required"`
	Name        string `json:"name" binding:"required,min=3,max=150"`
	Description string `json:"description"`
	Price       int64  `json:"price" binding:"required,min=4"`
	CoverUrl    string `json:"cover_url"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CoverUrl    string `json:"cover_url"`
	CreatedAt   time.Time
}

type ProductUpdateResponse struct {
	Slug        string `json:"slug"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CoverUrl    string `json:"cover_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
