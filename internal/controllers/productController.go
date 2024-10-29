package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dmitriy4565/VapeShop/internal/services/productService"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	productService *productService.ProductService
	validate       *validator.Validate
}

func NewProductController(productService *productService.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
		validate:       validator.New(),
	}
}

func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID продукта не указан", http.StatusBadRequest)
		return
	}

	product, err := c.productService.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product productService.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct, err := c.productService.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newProduct)
}

func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product productService.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.productService.UpdateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID продукта не указан", http.StatusBadRequest)
		return
	}

	err := c.productService.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
