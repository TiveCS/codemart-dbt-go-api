package controller

import (
	"log"
	"net/http"

	"github.com/TiveCS/codemart-dbt-go-api/model"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUsecase model.ProductUsecase
}

func NewProductController() *ProductController {
	return new(ProductController)
}

// RegisterProductUsecase implements model.ProductController
func (c *ProductController) RegisterProductUsecase(usecase model.ProductUsecase) {
	c.productUsecase = usecase
}

func (c *ProductController) GetAllProducts(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// Prepare response
	res := new(model.Response)
	data := new(model.GetAllProductsResponse)

	// Attempt to get all products
	products, err := c.productUsecase.GetAllProducts(ctx)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	// Construct payload response
	data.Count = len(products)
	data.Products = products

	// Send response
	res = model.NewResponse().WithMessage("Success get all products").WithData(data)
	return eCtx.JSON(http.StatusOK, res)
}

// CreateNewProduct implements model.ProductController
func (c *ProductController) CreateNewProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// Prepare response and request
	res := new(model.Response)
	req := new(model.CreateNewProductRequest)
	data := new(model.CreateNewProductResponse)

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
	data.ID = insertedId
	res = model.NewResponse().WithMessage("Success create new product").WithData(data)
	return eCtx.JSON(http.StatusOK, res)
}

func (c *ProductController) GetProductByID(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// Prepare response and request
	req := new(model.GetProductByIDRequest)
	res := new(model.Response)
	data := new(model.GetProductByIDResponse)

	// Get request params
	err := eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("Invalid request params")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	log.Println(req.ID)

	// Attempt to get product by id
	product, err := c.productUsecase.GetProductByID(ctx, req.ID)
	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	// Construct payload response
	data.Product = product

	// Send response
	res = model.NewResponse().WithMessage("Success get product by id").WithData(data)
	return eCtx.JSON(http.StatusOK, res)
}
