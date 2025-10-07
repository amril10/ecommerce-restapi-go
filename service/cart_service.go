package service

import (
	"errors"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/model"
	"github.com/amril10/rest-api-go/repository"
	"gorm.io/gorm"
)

type CartService interface {
	GetAll(userID int) ([]dto.CartResponse, error)
	CreateOrUpdateCart(req *dto.CartRequest, UserID int) (*dto.CartResponse, error)
	UpdateCart(id int, userID int, req *dto.UpdateCartRequest) (*dto.UpdateCartResponse, error)
	DeleteCart(id int) error
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(r repository.CartRepository) CartService {
	return &cartService{
		repo: r,
	}
}

func (s *cartService) GetAll(userID int) ([]dto.CartResponse, error) {
	carts, err := s.repo.GetAllCart(userID)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := make([]dto.CartResponse, len(carts))
	for i, c := range carts {
		response[i] = dto.CartResponse{
			ID:        c.ID,
			ProductID: c.ProductID,
			UserID:    c.UserID,
			Qty:       uint(c.Qty),
			Total:     c.Total,
		}
	}

	return response, nil
}

func (s *cartService) CreateOrUpdateCart(req *dto.CartRequest, userID int) (*dto.CartResponse, error) {
	product, err := s.repo.ExistProduct(req.ProductID)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Product not found"}
	}

	cart, err := s.repo.FindCartByUserAndProduct(userID, req.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCart := &model.Cart{
				ProductID: product.ID,
				UserID:    userID,
				Qty:       req.Qty,
				Total:     int(req.Qty) * int(product.Price),
			}

			if err := s.repo.CreateCart(newCart); err != nil {
				return nil, &errorhandler.InternalServerError{Message: err.Error()}
			}

			data := dto.CartResponse{
				ID:        newCart.ID,
				ProductID: newCart.ProductID,
				UserID:    newCart.UserID,
				Qty:       req.Qty,
				Total:     newCart.Total,
				CreatedAt: newCart.CreatedAt,
				UpdatedAt: newCart.UpdatedAt,
			}

			return &data, nil
		}
		return nil, err
	}

	cart.Qty += req.Qty
	cart.Total = int(cart.Qty) * int(product.Price)

	if err := s.repo.UpdateCart(cart); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	res := dto.CartResponse{
		ID:        cart.ID,
		ProductID: cart.ProductID,
		UserID:    cart.UserID,
		Qty:       uint(cart.Qty),
		Total:     cart.Total,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}

	return &res, nil
}

func (s *cartService) UpdateCart(id int, userID int, req *dto.UpdateCartRequest) (*dto.UpdateCartResponse, error) {
	cart, err := s.repo.FindCartByID(id, userID)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Cart not found"}
	}

	cart.Qty = req.Qty
	cart.Total = int(req.Qty) * int(cart.Product.Price)

	if err := s.repo.UpdateCart(cart); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	res := dto.UpdateCartResponse{
		ID:        cart.ID,
		ProductID: cart.ProductID,
		UserID:    cart.UserID,
		Qty:       uint(cart.Qty),
		Total:     cart.Total,
		IsSelect:  cart.IsSelect,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}

	return &res, nil
}

func (s *cartService) DeleteCart(id int) error {
	err := s.repo.DeleteCart(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Cart not found"}
	}

	return nil
}
