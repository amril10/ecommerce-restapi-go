package repository

import (
	"github.com/amril10/rest-api-go/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCart() ([]model.Cart, error)
	FindCartByUserAndProduct(ProductID, userID int) (*model.Cart, error)
	ExistProduct(ProductID int) (*model.Product, error)
	CreateCart(cart *model.Cart) error
	FindCartByID(id int, userID int) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	DeleteCart(id int) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) GetAllCart() ([]model.Cart, error) {
	var carts []model.Cart
	err := r.db.Find(&carts).Order("updated_at DESC").Error
	return carts, err
}

func (r *cartRepository) FindCartByUserAndProduct(userID, ProductID int) (*model.Cart, error) {
	var existCart model.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, ProductID).First(&existCart).Error
	if err != nil {
		return nil, err
	}

	return &existCart, nil
}

func (r *cartRepository) ExistProduct(ProductID int) (*model.Product, error) {
	var existProduct model.Product
	err := r.db.First(&existProduct, ProductID).Error
	if err != nil {
		return nil, err
	}

	return &existProduct, nil
}

func (r *cartRepository) CreateCart(cart *model.Cart) error {
	err := r.db.Create(cart).Error
	return err
}

func (r *cartRepository) FindCartByID(id int, userID int) (*model.Cart, error) {
	var cart model.Cart
	if err := r.db.Preload("Product").Where("id = ? AND user_id = ?", id, userID).First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepository) UpdateCart(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) DeleteCart(id int) error {
	var cart model.Cart
	if err := r.db.First(&cart, id).Error; err != nil {
		return err
	}

	return r.db.Delete(&cart).Error
}
