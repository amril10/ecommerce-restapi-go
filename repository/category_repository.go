package repository

import (
	"github.com/amril10/rest-api-go/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(req *model.Category) error
	GetAllCategory() ([]model.Category, error)
	GetCategoryBySlug(slug string) (*model.Category, error)
	UpdateCategory(id uint, category *model.Category) (*model.Category, error)
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAllCategory() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryBySlug(slug string) (*model.Category, error) {
	var categories model.Category
	err := r.db.Where("slug = ?", slug).First(&categories).Error
	return &categories, err
}

func (r *categoryRepository) CreateCategory(category *model.Category) error {
	err := r.db.Create(category).Error
	return err
}

func (r *categoryRepository) UpdateCategory(id uint, category *model.Category) (*model.Category, error) {
	var existing model.Category
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}

	existing.Slug = category.Slug
	existing.Nama = category.Nama

	if err := r.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	var category model.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}

	return r.db.Delete(&category).Error
}
