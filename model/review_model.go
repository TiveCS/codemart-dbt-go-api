package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

type ReviewRepository interface {
	CreateNewReview(ctx context.Context, review *Review) (*Review, error)
	GetReviewByID(ctx context.Context, id primitive.ObjectID) (*Review, error)
	GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID) ([]*Review, error)
	GetAllReviews(ctx context.Context) ([]*Review, error)
}

type ReviewUsecase interface {
	CreateNewReview(ctx context.Context, review *Review) (*Review, error)
	GetReviewByID(ctx context.Context, id string) (*Review, error)
	GetReviewsByProductID(ctx context.Context, productID string) ([]*Review, error)
	GetAllReviews(ctx context.Context) ([]*Review, error)

	RegisterReviewRepository(repo ReviewRepository)
}

type ReviewController interface {
	RegisterReviewUsecase(usecase ReviewUsecase)

	CreateNewReview(eCtx echo.Context) error
	GetReviewByID(eCtx echo.Context) error
	GetReviewsByProductID(eCtx echo.Context) error
	GetAllReviews(eCtx echo.Context) error
}

type CreateNewReviewRequest struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id" param:"product_id" form:"product_id"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
}

type GetAllReviewsResponse struct {
	Reviews []*Review `json:"reviews" bson:"reviews"`
	Count   int       `json:"count" bson:"count"`
}

type GetReviewByIDResponse struct {
	Review *Review `json:"review" bson:"review"`
}

type GetReviewsByProductIDResponse struct {
	Reviews []*Review `json:"reviews" bson:"reviews"`
	Count   int       `json:"count" bson:"count"`
}

func (m *CreateNewReviewRequest) ToReview() *Review {
	return &Review{
		ProductID: m.ProductID,
		Content:   m.Content,
		Author:    m.Author,
		CreatedAt: time.Now().Unix(),
	}
}
