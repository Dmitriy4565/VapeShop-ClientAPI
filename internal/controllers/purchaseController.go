package controllers

import (
	"net/http"

	"VapeShop-ClientAPI/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PurchaseController - контроллер для работы с покупками.
type PurchaseController struct {
	purchaseService services.PurchaseService
	validate        *validator.Validate
}

// NewPurchaseController - функция для создания нового контроллера покупок.
func NewPurchaseController(purchaseService services.PurchaseService) *PurchaseController {
	return &PurchaseController{
		purchaseService: purchaseService,
		validate:        validator.New(),
	}
}

// GetPurchasesHandler - обработчик запроса на получение всех покупок.
func (c *PurchaseController) GetPurchasesHandler(ctx *gin.Context) {
	purchases, err := c.purchaseService.GetAllPurchases(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, purchases)
}

// GetPurchaseByIDHandler - обработчик запроса на получение покупки по ID.
func (c *PurchaseController) GetPurchaseByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID покупки не указан"})
		return
	}

	purchase, err := c.purchaseService.GetPurchaseByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, purchase)
}

// CreatePurchaseHandler - обработчик запроса на создание новой покупки.
func (c *PurchaseController) CreatePurchaseHandler(ctx *gin.Context) {
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

	newPurchase, err := c.purchaseService.CreatePurchase(ctx, purchase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newPurchase)
}

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

// DeletePurchaseHandler - обработчик запроса на удаление покупки
func (c *PurchaseController) DeletePurchaseHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID покупки не указан"})
		return
	}

	err := c.purchaseService.DeletePurchase(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Покупка успешно удалена"})
}
