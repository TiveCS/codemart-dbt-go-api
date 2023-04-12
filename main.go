package main

import (
	"github.com/TiveCS/codemart-dbt-go-api/controller"
	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/TiveCS/codemart-dbt-go-api/repository"
	"github.com/TiveCS/codemart-dbt-go-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db.ConnectMongo()

	productRepository := repository.NewProductRepository()
	productUsecase := usecase.NewProductUsecase()
	productUsecase.RegisterProductRepository(productRepository)
	productController := controller.NewProductController()
	productController.RegisterProductUsecase(productUsecase)

	reviewRepository := repository.NewReviewRepository()
	reviewUsecase := usecase.NewReviewUsecase()
	reviewUsecase.RegisterReviewRepository(reviewRepository)
	reviewController := controller.NewReviewController()
	reviewController.RegisterReviewUsecase(reviewUsecase)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/products", productController.CreateNewProduct)
	e.GET("/products", productController.GetAllProducts)
	e.GET("/products/:id", productController.GetProductByID)

	e.POST("/products/:product_id/reviews", reviewController.CreateNewReview)
	e.GET("/products/:product_id/reviews", reviewController.GetReviewsByProductID)
	e.GET("/reviews", reviewController.GetAllReviews)

	e.Logger.Fatal(e.Start(":1323"))
}
