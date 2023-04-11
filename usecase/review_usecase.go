package usecase

import (
	"context"

	"github.com/TiveCS/codemart-dbt-go-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reviewUsecase struct {
	reviewRepository model.ReviewRepository
}

func NewReviewUsecase() model.ReviewUsecase {
	return new(reviewUsecase)
}

// CreateNewReview implements model.ReviewUsecase
func (u *reviewUsecase) CreateNewReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	return u.reviewRepository.CreateNewReview(ctx, review)
}

// GetAllReviews implements model.ReviewUsecase
func (u *reviewUsecase) GetAllReviews(ctx context.Context) ([]*model.Review, error) {
	return u.reviewRepository.GetAllReviews(ctx)
}

// GetReviewsByProductID implements model.ReviewUsecase
func (u *reviewUsecase) GetReviewsByProductID(ctx context.Context, productID string) ([]*model.Review, error) {
	productObjID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	return u.reviewRepository.GetReviewsByProductID(ctx, productObjID)
}

// RegisterReviewRepository implements model.ReviewUsecase
func (u *reviewUsecase) RegisterReviewRepository(repo model.ReviewRepository) {
	u.reviewRepository = repo
}
