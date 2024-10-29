package main

import (
	"net/http"

	"github.com/Dmitriy4565/VapeShop/internal/db"
	"github.com/Dmitriy4565/VapeShop/internal/services"
	"github.com/gin-gonic/gin" // Используем Gin для HTTP-обработки
)

type Server struct {
	router          *gin.Engine
	categoryService services.CategoryService
}

func NewServer(db *db.DB) *Server {
	router := gin.Default()

	categoryService := services.NewCategoryService(db)

	categoryController := NewCategoryController(categoryService)

	router.GET("/categories", categoryController.GetCategoriesHandler)

	return &Server{
		router:          router,
		categoryService: categoryService,
	}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
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

