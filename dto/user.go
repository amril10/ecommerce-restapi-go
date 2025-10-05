package dto

import "time"

type ProfileResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

type ProfileUpdateRequest struct {
	Name string `json:"name" binding:"required,min=3,max=150"`
}

type ProfileUpdateResponse struct {
	Name      string `json:"name"`
	UpdatedAt time.Time
}
