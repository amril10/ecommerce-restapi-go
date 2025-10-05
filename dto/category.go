package dto

import "time"

type CategoryRequest struct {
	Slug string `json:"slug" binding:"required,min=3,max=100"`
	Nama string `json:"nama" binding:"required,min=3,max=100"`
}

type CategoryResponse struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Nama      string `json:"nama"`
	CreatedAt time.Time
}

type UpdateCategoryResponse struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Nama      string `json:"nama"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
