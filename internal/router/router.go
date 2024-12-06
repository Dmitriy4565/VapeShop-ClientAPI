package router

import (
	"VapeShop-ClientAPI/internal/config"
	"VapeShop-ClientAPI/internal/controllers"
	"VapeShop-ClientAPI/internal/db"
	"VapeShop-ClientAPI/internal/middleware"
	"VapeShop-ClientAPI/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	categoryService services.CategoryService
	productService  services.ProductService
	purchaseService services.PurchaseService
}

func NewServer(cfg *config.Config, db *db.DB) *Server {
	router := gin.Default()

	// Middleware
	router.Use(middleware.Cors())
	router.Use(middleware.AuthMiddleware())

	categoryService := services.NewCategoryService(db.DB)
	productService := services.NewProductService(db.DB)
	purchaseService := services.NewPurchaseService(db.DB)

	// Controllers
	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	purchaseController := controllers.NewPurchaseController(purchaseService)

	RegisterRoutes(router, categoryController, productController, purchaseController)

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

func RegisterRoutes(router *gin.Engine, categoryController *controllers.CategoryController, productController *controllers.ProductController, purchaseController *controllers.PurchaseController) {
	v1 := router.Group("/api/v1")

	v1.GET("/categories", categoryController.GetCategoriesHandler)
	v1.POST("/categories", categoryController.CreateCategoriesHandler)
	v1.GET("/categories/:id", categoryController.GetCategoryByIDHandler)
	v1.PUT("/categories/:id", categoryController.UpdateCategoryHandler)
	v1.DELETE("/categories/:id", categoryController.DeleteCategoryHandler)

	v1.GET("/products", productController.GetProductsHandler)
	v1.POST("/products", productController.CreateProductHandler)
	v1.GET("/products/:id", productController.GetProductByIDHandler)
	v1.PUT("/products/:id", productController.UpdateProductHandler)
	v1.DELETE("/products/:id", productController.DeleteProductHandler)

	v1.GET("/purchases", purchaseController.GetPurchasesHandler)
	v1.POST("/purchases", purchaseController.CreatePurchaseHandler)
	v1.GET("/purchases/:id", purchaseController.GetPurchaseByIDHandler)
	v1.PUT("/purchases/:id", purchaseController.UpdatePurchaseHandler)
	v1.DELETE("/purchases/:id", purchaseController.DeletePurchaseHandler)
}
