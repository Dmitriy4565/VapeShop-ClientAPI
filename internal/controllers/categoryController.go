package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dmitriy4565/VapeShop/internal/services/categoryService"
	"github.com/go-playground/validator/v10"
)

type CategoryController struct {
	categoryService *categoryService.CategoryService
	validate        *validator.Validate
}

func NewCategoryController(categoryService *categoryService.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
		validate:        validator.New(),
	}
}

func (c *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (c *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category categoryService.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory, err := c.categoryService.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCategory)
}

func (c *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category categoryService.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.categoryService.UpdateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID категории не указан", http.StatusBadRequest)
		return
	}

	err := c.categoryService.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
