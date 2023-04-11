package repository

import (
	"context"

	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/TiveCS/codemart-dbt-go-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	allReviews          = []*model.Review{}
	allReviewsHasUpdate = true
)

type reviewRepository struct {
	productCollection *mongo.Collection
}

func NewReviewRepository() model.ReviewRepository {
	return &reviewRepository{
		productCollection: db.GetMongoCollection("products"),
	}
}

// CreateNewReview implements model.ReviewRepository
func (r *reviewRepository) CreateNewReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	_, err := r.productCollection.UpdateOne(ctx, bson.M{"_id": review.ProductID}, bson.M{"$push": bson.M{"reviews": review}})
	if err != nil {
		return nil, err
	}

	allReviewsHasUpdate = true

	return review, err
}

// GetAllReviews implements model.ReviewRepository
func (r *reviewRepository) GetAllReviews(ctx context.Context) ([]*model.Review, error) {
	if !allReviewsHasUpdate {
		return allReviews, nil
	}

	options := options.Find().SetProjection(bson.M{"reviews": 1})
	cursor, err := r.productCollection.Find(ctx, bson.M{}, options)

	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for cursor.Next(ctx) {
		var product *model.Product
		err := cursor.Decode(&product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	var returnReviews []*model.Review
	for _, product := range products {
		returnReviews = append(returnReviews, product.Reviews...)
	}

	if allReviewsHasUpdate {
		allReviews = returnReviews
		allReviewsHasUpdate = false
	}

	return returnReviews, nil
}

// GetReviewByID implements model.ReviewRepository
func (r *reviewRepository) GetReviewByID(ctx context.Context, id primitive.ObjectID) (*model.Review, error) {
	result := r.productCollection.FindOne(ctx, bson.M{"_id": id})

	var review *model.Review
	err := result.Decode(&review)

	if err != nil {
		return nil, err
	}

	return review, nil
}

// GetReviewsByProductID implements model.ReviewRepository
func (r *reviewRepository) GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID) ([]*model.Review, error) {
	productResult := r.productCollection.FindOne(ctx, bson.M{"_id": productID})

	if productResult.Err() != nil {
		return nil, productResult.Err()
	}

	var product *model.Product
	err := productResult.Decode(&product)

	if err != nil {
		return nil, err
	}

	var reviews []*model.Review
	reviews = append(reviews, product.Reviews...)

	return reviews, nil
}
