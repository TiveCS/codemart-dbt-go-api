package repository

import (
	"context"

	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/TiveCS/codemart-dbt-go-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	productCollection *mongo.Collection
}

func NewProductRepository() model.ProductRepository {
	return &productRepository{
		productCollection: db.GetMongoCollection("products"),
	}
}

// Create implements model.ProductRepository
func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	r.productCollection.InsertOne(ctx, product)
	return nil
}

// FindAll implements model.ProductRepository
func (r *productRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	result, err := r.productCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for result.Next(ctx) {
		var product *model.Product
		err := result.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

// FindByID implements model.ProductRepository
func (r *productRepository) FindByID(ctx context.Context, id int) (*model.Product, error) {
	result := r.productCollection.FindOne(ctx, map[string]interface{}{
		"id": id,
	})

	var product *model.Product
	err := result.Decode(&product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
