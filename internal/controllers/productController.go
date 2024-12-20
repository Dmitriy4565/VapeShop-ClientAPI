package controllers

import (
	"net/http"
	"strconv"

	"VapeShop-ClientAPI/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ProductController - контроллер для работы с продуктами.
type ProductController struct {
	productService services.ProductService
	validate       *validator.Validate
}

// NewProductController - функция для создания нового контроллера продуктов.
func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
		validate:       validator.New(),
	}
}

// GetProductsHandler - обработчик запроса на получение всех продуктов.
func (c *ProductController) GetProductsHandler(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// GetProductByIDHandler - обработчик запроса на получение продукта по ID.
func (c *ProductController) GetProductByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	product, err := c.productService.GetProductByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// CreateProductHandler - обработчик запроса на создание нового продукта.
func (c *ProductController) CreateProductHandler(ctx *gin.Context) {
	var product services.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.validate.Struct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := c.productService.CreateProduct(ctx, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newProduct)
}

func (c *ProductController) UpdateProductHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ID"})
		return
	}

	var product services.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = id // Устанавливаем ID для обновления

	updatedProduct, err := c.productService.UpdateProduct(ctx.Request.Context(), product)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

// DeleteProductHandler - обработчик запроса на удаление продукта.
func (c *ProductController) DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID продукта не указан"})
		return
	}

	err := c.productService.DeleteProduct(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Продукт успешно удален"})
}
