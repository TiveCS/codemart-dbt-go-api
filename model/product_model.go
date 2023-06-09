package model

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Owner       string             `bson:"owner" json:"owner"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       int64              `bson:"price" json:"price"`
	CoverURL    string             `bson:"cover_url" json:"cover_url"`
	Reviews     []*Review          `bson:"reviews" json:"reviews"`
}

type ProductUsecase interface {
	CreateNewProduct(ctx context.Context, product *Product) (primitive.ObjectID, error)
	GetAllProducts(ctx context.Context) ([]*Product, error)
	GetProductByID(ctx context.Context, id string) (*Product, error)

	RegisterProductRepository(repo ProductRepository)
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*Product, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Product, error)
	Create(ctx context.Context, product *Product) (primitive.ObjectID, error)
}

type ProductController interface {
	CreateNewProduct(eCtx echo.Context) error
	GetAllProducts(eCtx echo.Context) error
	GetProductByID(eCtx echo.Context) error

	RegisterProductUsecase(usecase ProductUsecase)
}

type CreateNewProductRequest struct {
	Owner       string `json:"owner"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CoverURL    string `json:"cover_url"`
}

type CreateNewProductResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type GetAllProductsResponse struct {
	Products []*Product `json:"products"`
	Count    int        `json:"count"`
}

type GetProductByIDRequest struct {
	ID string `json:"id" query:"id" param:"id" form:"id" bson:"_id"`
}

type GetProductByIDResponse struct {
	Product *Product `json:"product"`
}

func (m *CreateNewProductRequest) ToProduct() *Product {
	return &Product{
		ID:          primitive.NewObjectID(),
		Owner:       m.Owner,
		Title:       m.Title,
		Description: m.Description,
		Price:       m.Price,
		CoverURL:    m.CoverURL,
		Reviews:     []*Review{},
	}
}
