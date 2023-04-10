package controller

import (
	"log"
	"net/http"

	"github.com/TiveCS/codemart-dbt-go-api/model"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productUsecase model.ProductUsecase
}

func NewProductController() model.ProductController {
	return new(productController)
}

// RegisterProductUsecase implements model.ProductController
func (c *productController) RegisterProductUsecase(usecase model.ProductUsecase) {
	c.productUsecase = usecase
}

func (c *productController) GetAllProducts(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// Prepare response
	res := new(model.Response)

	// Attempt to get all products
	products, err := c.productUsecase.GetAllProducts(ctx)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		log.Println("Error: ", err)
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	// Send response
	res = model.NewResponse().WithMessage("Success get all products").WithData(map[string]interface{}{
		"counts":   len(products),
		"products": products,
	})
	return eCtx.JSON(http.StatusOK, res)
}

// CreateNewProduct implements model.ProductController
func (c *productController) CreateNewProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// Prepare response and request
	res := new(model.Response)
	req := new(model.CreateNewProductRequest)

	// Validate request body
	err := eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("Invalid request body")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	// Convert request to product model
	product := req.ToProduct()

	// Attempt to persist the payload
	insertedId, err := c.productUsecase.CreateNewProduct(ctx, product)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	// Send response
	res = model.NewResponse().WithMessage("Success create new product").WithData(map[string]interface{}{
		"id": insertedId,
	})
	return eCtx.JSON(http.StatusOK, res)
}
