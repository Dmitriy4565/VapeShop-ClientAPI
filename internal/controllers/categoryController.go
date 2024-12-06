package controllers

import (
	"net/http"
	"strconv"

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

func (c *CategoryController) GetCategoryByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id") // Получаем ID из параметров URL

	category, err := c.categoryService.GetCategoryByID(ctx, id) // Вызов сервиса для получения категории
	if err != nil {
		if err.Error() == "категория не найдена" { // Обработка ошибки, если категория не найдена
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Обработка других ошибок
		return
	}

	ctx.JSON(http.StatusOK, category) // Возвращаем найденную категорию
}

func (c *CategoryController) CreateCategoriesHandler(ctx *gin.Context) {
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

// UpdateCategoryHandler - обработчик запроса на обновление категории.
func (c *CategoryController) UpdateCategoryHandler(ctx *gin.Context) {
	var category services.Category

	// Извлечение ID из параметров URL
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый ID"})
		return
	}

	// Привязка JSON к структуре Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID для обновления
	category.ID = id

	// Валидация структуры Category
	err = c.validate.Struct(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление категории через сервис
	updatedCategory, err := c.categoryService.UpdateCategory(ctx.Request.Context(), category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправляем обновлённую категорию
	ctx.JSON(http.StatusOK, updatedCategory)
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
