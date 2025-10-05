package service

import (
	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/model"
	"github.com/amril10/rest-api-go/repository"
)

type CategoryService interface {
	CreateCategory(req *dto.CategoryRequest) (*dto.CategoryResponse, error)
	GetAllCategory() ([]dto.CategoryResponse, error)
	GetCategoryBySlug(slug string) (*dto.CategoryResponse, error)
	UpdateCategory(id uint, req *dto.CategoryRequest) (*dto.UpdateCategoryResponse, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: r,
	}
}

func (s *categoryService) CreateCategory(req *dto.CategoryRequest) (*dto.CategoryResponse, error) {
	var data dto.CategoryResponse

	category := model.Category{
		Slug: req.Slug,
		Nama: req.Nama,
	}

	if err := s.repository.CreateCategory(&category); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.CategoryResponse{
		ID:        category.ID,
		Slug:      category.Slug,
		Nama:      category.Nama,
		CreatedAt: category.CreatedAt,
	}

	return &data, nil
}

func (s *categoryService) GetAllCategory() ([]dto.CategoryResponse, error) {

	categories, err := s.repository.GetAllCategory()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := make([]dto.CategoryResponse, len(categories))
	for i, c := range categories {
		response[i] = dto.CategoryResponse{
			ID:        c.ID,
			Slug:      c.Slug,
			Nama:      c.Nama,
			CreatedAt: c.CreatedAt,
		}
	}

	return response, nil
}

func (s *categoryService) GetCategoryBySlug(slug string) (*dto.CategoryResponse, error) {
	category, err := s.repository.GetCategoryBySlug(slug)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Category not found"}
	}

	response := &dto.CategoryResponse{
		ID:        category.ID,
		Slug:      category.Slug,
		Nama:      category.Nama,
		CreatedAt: category.CreatedAt,
	}

	return response, nil
}

func (s *categoryService) UpdateCategory(id uint, req *dto.CategoryRequest) (*dto.UpdateCategoryResponse, error) {
	category := &model.Category{
		Slug: req.Slug,
		Nama: req.Nama,
	}

	updated, err := s.repository.UpdateCategory(id, category)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := &dto.UpdateCategoryResponse{
		ID:        updated.ID,
		Slug:      updated.Slug,
		Nama:      updated.Nama,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}

	return response, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	err := s.repository.DeleteCategory(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Category not found"}
	}

	return nil
}
