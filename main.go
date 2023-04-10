package main

import (
	"net/http"

	"github.com/TiveCS/codemart-dbt-go-api/controller"
	"github.com/TiveCS/codemart-dbt-go-api/db"
	"github.com/TiveCS/codemart-dbt-go-api/repository"
	"github.com/TiveCS/codemart-dbt-go-api/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db.ConnectMongo()

	productRepository := repository.NewProductRepository()
	productUsecase := usecase.NewProductUsecase()
	productUsecase.RegisterProductRepository(productRepository)
	productController := controller.NewProductController()
	productController.RegisterProductUsecase(productUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, Shadow!",
		})
	})

	e.POST("/products", productController.CreateNewProduct)
	e.GET("/products", productController.GetAllProducts)

	e.Logger.Fatal(e.Start(":1323"))
}
