package repository

import (
	"github.com/amril10/rest-api-go/model"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(userID int) (*model.User, error)
	FindProfileByID(userID int) (*model.User, error)
	UpdateProfile(profile *model.User) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) GetProfile(userID int) (*model.User, error) {
	var profile model.User
	err := r.db.Where("id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) FindProfileByID(userID int) (*model.User, error) {
	var profile model.User
	err := r.db.Where("id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepository) UpdateProfile(profile *model.User) error {
	return r.db.Save(profile).Error
}
