package main

import (
	"net/http"

	"VapeShop-ClientAPI/internal/config"
	"VapeShop-ClientAPI/internal/controllers"
	"VapeShop-ClientAPI/internal/db"
	"VapeShop-ClientAPI/internal/middleware"
	"VapeShop-ClientAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	categoryService services.CategoryService
	productService  services.ProductService
	purchaseService services.PurchaseService
}

func NewServer(db *db.DB, cfg *config.Config) *Server {
	router := gin.Default()

	router.Use(middleware.Cors())

	categoryService := services.NewCategoryService(db.DB)
	productService := services.NewProductService(db.DB)
	purchaseService := services.NewPurchaseService(db.DB)

	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	purchaseController := controllers.NewPurchaseController(purchaseService)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/categories", categoryController.GetCategoriesHandler)
		v1.GET("/products", productController.GetProductsHandler)
		v1.POST("/purchases", purchaseController.CreatePurchaseHandler)
	}

	return &Server{
		router:          router,
		categoryService: categoryService,
		productService:  productService,
		purchaseService: purchaseService,
	}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDB(cfg.DB)
	if err != nil {
		panic(err)
	}

	server := NewServer(db, cfg)
	if err := server.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
