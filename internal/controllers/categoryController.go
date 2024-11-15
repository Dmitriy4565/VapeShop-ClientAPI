package controllers

import (
	"net/http"

	"VapeShop-ClientAPI/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryController struct {
	categoryService services.CategoryService
	validate        *validator.Validate
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
		validate:        validator.New(),
	}
}

func (c *CategoryController) GetCategoriesHandler(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) CreateCategoryHandler(ctx *gin.Context) {
	var category services.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.validate.Struct(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory, err := c.categoryService.CreateCategory(ctx, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newCategory)
}

// UpdatePurchaseHandler - обработчик запроса на обновление покупки.
func (c *PurchaseController) UpdatePurchaseHandler(ctx *gin.Context) {
	var purchase services.Purchase
	if err := ctx.ShouldBindJSON(&purchase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.validate.Struct(purchase)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPurchase, err := c.purchaseService.UpdatePurchase(ctx, purchase) // Изменено: теперь получаем обновлённую покупку
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedPurchase) // Отправляем обновлённую покупку
}

func (c *CategoryController) DeleteCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	err := c.categoryService.DeleteCategory(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
