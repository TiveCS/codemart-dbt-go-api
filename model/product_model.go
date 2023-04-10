package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int    `json:"id"`
	Owner       int    `json:"owner"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CoverURL    string `json:"cover_url"`
}

type ProductUsecase interface {
	CreateNewProduct(ctx context.Context) error
	GetAllProducts(ctx context.Context) ([]*Product, error)
	GetProductByID(ctx context.Context) (*Product, error)
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*Product, error)
	FindByID(ctx context.Context, id int) (*Product, error)
	Create(ctx context.Context, product *Product) error
}

type ProductController interface {
	CreateNewProduct(c echo.Context) error
}
