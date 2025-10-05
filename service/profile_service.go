package service

import (
	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/repository"
)

type ProfileService interface {
	GetProfile(userID int) (*dto.ProfileResponse, error)
	UpdateProfile(userID int, req *dto.ProfileUpdateRequest) (*dto.ProfileUpdateResponse, error)
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(r repository.ProfileRepository) *profileService {
	return &profileService{
		repo: r,
	}
}

func (s *profileService) GetProfile(userID int) (*dto.ProfileResponse, error) {
	profile, err := s.repo.GetProfile(userID)
	if err != nil {
		return nil, &errorhandler.ForbiddenError{Message: "Invalid access"}
	}

	response := dto.ProfileResponse{
		ID:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		CreatedAt: profile.CreatedAt,
	}

	return &response, nil
}

func (s *profileService) UpdateProfile(userID int, req *dto.ProfileUpdateRequest) (*dto.ProfileUpdateResponse, error) {
	profile, err := s.repo.FindProfileByID(userID)
	if err != nil {
		return nil, &errorhandler.ForbiddenError{Message: "Invalid access"}
	}

	profile.Name = req.Name

	if err := s.repo.UpdateProfile(profile); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := dto.ProfileUpdateResponse{
		Name:      profile.Name,
		UpdatedAt: profile.UpdatedAt,
	}

	return &response, nil
}
