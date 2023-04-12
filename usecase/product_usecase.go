package usecase

import (
	"context"

	"github.com/TiveCS/codemart-dbt-go-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productUsecase struct {
	productRepository model.ProductRepository
}

func NewProductUsecase() model.ProductUsecase {
	return new(productUsecase)
}

// RegisterProductRepository implements model.ProductUsecase
func (u *productUsecase) RegisterProductRepository(repo model.ProductRepository) {
	u.productRepository = repo
}

// CreateNewProduct implements model.ProductUsecase
func (u *productUsecase) CreateNewProduct(ctx context.Context, product *model.Product) (primitive.ObjectID, error) {
	return u.productRepository.Create(ctx, product)
}

// GetAllProducts implements model.ProductUsecase
func (u *productUsecase) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	return u.productRepository.FindAll(ctx)
}

// GetProductByID implements model.ProductUsecase
func (u *productUsecase) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return u.productRepository.FindByID(ctx, objID)
}
