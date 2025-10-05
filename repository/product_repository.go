package repository

import (
	"github.com/amril10/rest-api-go/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
	GetProductBySlug(slug string) (*model.Product, error)
	CreateProduct(req *model.Product) error
	UpdateProduct(id int, req *model.Product) (*model.Product, error)
	CheckCategoryID(id int) (bool, error)
	DeleteProduct(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Order("updated_at DESC").Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductBySlug(slug string) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("slug = ?", slug).First(&product).Error
	return &product, err
}

func (r *productRepository) CheckCategoryID(id int) (bool, error) {
	var ExistCategory int64
	err := r.db.Model(&model.Category{}).Where("id = ?", id).Count(&ExistCategory).Error
	if err != nil {
		return false, err
	}
	return ExistCategory > 0, nil
}

func (r *productRepository) CreateProduct(products *model.Product) error {
	err := r.db.Create(products).Error
	return err
}

func (r *productRepository) UpdateProduct(id int, req *model.Product) (*model.Product, error) {
	var existing model.Product
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}

	existing.Slug = req.Slug
	existing.Name = req.Name
	existing.Description = req.Description
	existing.Price = req.Price
	existing.CoverUrl = req.CoverUrl

	if err := r.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}

	return r.db.Delete(&product).Error
}
