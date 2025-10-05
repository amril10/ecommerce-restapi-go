package dto

type RegisterRequest struct {
	RoleID               int    `json:"role_id" binding:"required,min=1,max=3"`
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
