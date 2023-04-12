package repository

import (
	"context"

	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/TiveCS/codemart-dbt-go-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	allProductsCache     = []*model.Product{}
	allProductsHasUpdate = true

	productsCache          = map[primitive.ObjectID]*model.Product{}
	productsCacheHasUpdate = map[primitive.ObjectID]bool{}
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
func (r *productRepository) Create(ctx context.Context, product *model.Product) (primitive.ObjectID, error) {
	result, err := r.productCollection.InsertOne(ctx, product)

	if err != nil {
		return primitive.NilObjectID, err
	}

	allProductsHasUpdate = true

	return result.InsertedID.(primitive.ObjectID), nil
}

// FindAll implements model.ProductRepository
func (r *productRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	if !allProductsHasUpdate {
		return allProductsCache, nil
	}

	result, err := r.productCollection.Find(ctx, bson.M{})
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

	allProductsCache = products
	allProductsHasUpdate = false

	return products, nil
}

// FindByID implements model.ProductRepository
func (r *productRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Product, error) {
	cache := productsCache[id]
	hasUpdate := productsCacheHasUpdate[id]
	if cache != nil && !hasUpdate {
		return cache, nil
	}

	result := r.productCollection.FindOne(ctx, bson.M{"_id": id})

	var product *model.Product
	err := result.Decode(&product)
	if err != nil {
		return nil, err
	}

	productsCache[id] = product
	productsCacheHasUpdate[id] = false

	return product, nil
}
