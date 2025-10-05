package service

import (
	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/model"
	"github.com/amril10/rest-api-go/repository"
)

type ProductService interface {
	GetAllProduct() ([]dto.ProductResponse, error)
	GetProductBySlug(slug string) (*dto.ProductResponse, error)
	CreateProduct(req *dto.ProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(id int, req *dto.ProductRequest) (*dto.ProductUpdateResponse, error)
	DeleteProduct(id int) error
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		repository: r,
	}
}

func (s *productService) GetAllProduct() ([]dto.ProductResponse, error) {
	products, err := s.repository.GetAllProduct()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := make([]dto.ProductResponse, len(products))
	for i, c := range products {
		response[i] = dto.ProductResponse{
			ID:          c.ID,
			Slug:        c.Slug,
			CategoryID:  c.CategoryID,
			Name:        c.Name,
			Description: c.Description,
			Price:       c.Price,
			CoverUrl:    c.CoverUrl,
			CreatedAt:   c.CreatedAt,
		}
	}

	return response, nil
}

func (s *productService) GetProductBySlug(slug string) (*dto.ProductResponse, error) {
	product, err := s.repository.GetProductBySlug(slug)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Product not found"}
	}

	response := dto.ProductResponse{
		ID:          product.ID,
		Slug:        product.Slug,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CoverUrl:    product.CoverUrl,
		CreatedAt:   product.CreatedAt,
	}

	return &response, nil
}

func (s *productService) CreateProduct(req *dto.ProductRequest) (*dto.ProductResponse, error) {
	exist, err := s.repository.CheckCategoryID(req.CategoryID)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	if !exist {
		return nil, &errorhandler.NotFoundError{Message: "Category not found"}
	}

	products := model.Product{
		Slug:        req.Slug,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       int64(req.Price),
		CoverUrl:    req.CoverUrl,
	}

	if err := s.repository.CreateProduct(&products); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data := &dto.ProductResponse{
		ID:          products.ID,
		Slug:        products.Slug,
		CategoryID:  products.CategoryID,
		Name:        products.Name,
		Description: products.Description,
		Price:       products.Price,
		CoverUrl:    products.CoverUrl,
		CreatedAt:   products.CreatedAt,
	}

	return data, nil
}

func (s *productService) UpdateProduct(id int, req *dto.ProductRequest) (*dto.ProductUpdateResponse, error) {
	product := model.Product{
		Slug:        req.Slug,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CoverUrl:    req.CoverUrl,
	}

	updated, err := s.repository.UpdateProduct(id, &product)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	response := &dto.ProductUpdateResponse{
		Slug:        updated.Slug,
		CategoryID:  updated.CategoryID,
		Name:        updated.Name,
		Description: updated.Description,
		Price:       updated.Price,
		CoverUrl:    updated.CoverUrl,
		CreatedAt:   updated.CreatedAt,
		UpdatedAt:   updated.UpdatedAt,
	}

	return response, nil
}

func (s *productService) DeleteProduct(id int) error {
	err := s.repository.DeleteProduct(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Product not found"}
	}

	return nil
}
