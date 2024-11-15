package middleware

import (
	"net/http"

	"VapeShop-ClientAPI/internal/config"
	"VapeShop-ClientAPI/internal/controllers"
	"VapeShop-ClientAPI/internal/db"
	"VapeShop-ClientAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	categoryService services.CategoryService
	productService  services.ProductService
	purchaseService services.PurchaseService
}

// Middleware for CORS
func CorsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			return
		}
		c.Next()
	}
}

func NewServer(db *db.DB, cfg *config.Config) *Server {
	router := gin.Default()

	router.Use(CorsAuth())       // Use the renamed CORS middleware
	router.Use(AuthMiddleware()) // Use JWT authentication middleware

	// Create services -- CORRECTED: Use db directly, not db.Pool
	categoryService := services.NewCategoryService(db.DB)
	productService := services.NewProductService(db.DB)
	purchaseService := services.NewPurchaseService(db.DB)

	// Create controllers
	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	purchaseController := controllers.NewPurchaseController(purchaseService)

	// Set up routes -- CORRECTED: Use Gin's HandlerFunc correctly
	v1 := router.Group("/api/v1")
	{
		v1.GET("/categories", categoryController.GetCategoriesHandler)
		v1.GET("/products", productController.GetProductsHandler)
		v1.POST("/purchases", purchaseController.CreatePurchaseHandler)
		v1.GET("/purchases", purchaseController.GetPurchasesHandler)
		v1.GET("/purchases/:id", purchaseController.GetPurchaseByIDHandler)
		v1.PUT("/purchases/:id", purchaseController.UpdatePurchaseHandler)
		v1.DELETE("/purchases/:id", purchaseController.DeletePurchaseHandler)
	}

	return &Server{
		router:          router,
		categoryService: categoryService,
		productService:  productService,
		purchaseService: purchaseService,
	}
}
