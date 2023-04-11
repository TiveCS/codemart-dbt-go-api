package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TiveCS/codemart-dbt-go-api/model"
	"github.com/labstack/echo/v4"
)

type reviewController struct {
	reviewUsecase model.ReviewUsecase
}

func NewReviewController() model.ReviewController {
	return new(reviewController)
}

// CreateNewReview implements model.ReviewController
func (c *reviewController) CreateNewReview(eCtx echo.Context) error {
	processStart := time.Now().UnixNano()
	ctx := eCtx.Request().Context()

	req := new(model.CreateNewReviewRequest)
	res := new(model.Response)
	insertAmount := eCtx.QueryParam("amount")
	if insertAmount == "" {
		insertAmount = "1"
	}

	amount, err := strconv.Atoi(insertAmount)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	for i := 0; i < amount; i++ {
		review := req.ToReview()
		_, err = c.reviewUsecase.CreateNewReview(ctx, review)
		if err != nil {
			res = model.NewResponse().WithMessage(err.Error())
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
	}

	processEnd := time.Now().UnixNano()

	res = model.NewResponse().WithMessage("Success create new review").WithProcessTime(processStart, processEnd)
	return eCtx.JSON(http.StatusOK, res)
}

// GetAllReviews implements model.ReviewController
func (c *reviewController) GetAllReviews(eCtx echo.Context) error {
	processStart := time.Now().UnixNano()
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	data := new(model.GetAllReviewsResponse)

	reviews, err := c.reviewUsecase.GetAllReviews(ctx)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	data.Count = len(reviews)
	data.Reviews = reviews

	processEnd := time.Now().UnixNano()
	res = model.NewResponse().WithMessage("Success get all reviews").WithData(data).WithProcessTime(processStart, processEnd)
	return eCtx.JSON(http.StatusOK, res)
}

// GetReviewByID implements model.ReviewController
func (c *reviewController) GetReviewByID(eCtx echo.Context) error {
	processStart := time.Now().UnixNano()
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	data := new(model.GetReviewByIDResponse)

	id := eCtx.Param("id")
	review, err := c.reviewUsecase.GetReviewByID(ctx, id)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	data.Review = review

	processEnd := time.Now().UnixNano()
	res = model.NewResponse().WithMessage("Success get review by id").WithData(data).WithProcessTime(processStart, processEnd)
	return eCtx.JSON(http.StatusOK, res)
}

// GetReviewsByProductID implements model.ReviewController
func (c *reviewController) GetReviewsByProductID(eCtx echo.Context) error {
	processStart := time.Now().UnixNano()
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	data := new(model.GetReviewsByProductIDResponse)

	productID := eCtx.Param("product_id")
	reviews, err := c.reviewUsecase.GetReviewsByProductID(ctx, productID)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	data.Count = len(reviews)
	data.Reviews = reviews

	processEnd := time.Now().UnixNano()
	res = model.NewResponse().WithMessage("Success get reviews by product id").WithData(data).WithProcessTime(processStart, processEnd)
	return eCtx.JSON(http.StatusOK, res)
}

// RegisterReviewUsecase implements model.ReviewController
func (c *reviewController) RegisterReviewUsecase(usecase model.ReviewUsecase) {
	c.reviewUsecase = usecase
}
