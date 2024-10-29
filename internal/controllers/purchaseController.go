package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dmitriy4565/VapeShop/internal/services/purchaseService"
	"github.com/go-playground/validator/v10"
)

type PurchaseController struct {
	purchaseService *purchaseService.PurchaseService
	validate        *validator.Validate
}

func NewPurchaseController(purchaseService *purchaseService.PurchaseService) *PurchaseController {
	return &PurchaseController{
		purchaseService: purchaseService,
		validate:        validator.New(),
	}
}

func (c *PurchaseController) GetPurchasesHandler(w http.ResponseWriter, r *http.Request) {
	purchases, err := c.purchaseService.GetAllPurchases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(purchases)
}

func (c *PurchaseController) GetPurchaseByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID покупки не указан", http.StatusBadRequest)
		return
	}

	purchase, err := c.purchaseService.GetPurchaseByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(purchase)
}

func (c *PurchaseController) CreatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	var purchase purchaseService.Purchase
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPurchase, err := c.purchaseService.CreatePurchase(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPurchase)
}

func (c *PurchaseController) UpdatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	var purchase purchaseService.Purchase
	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.purchaseService.UpdatePurchase(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *PurchaseController) DeletePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID покупки не указан", http.StatusBadRequest)
		return
	}

	err := c.purchaseService.DeletePurchase(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
