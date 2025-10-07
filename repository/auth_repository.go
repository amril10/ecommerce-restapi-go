package repository

import (
	"github.com/amril10/rest-api-go/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(req *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *model.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *authRepository) EmailExist(email string) bool {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *authRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}
