package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Product - структура, представляющая продукт.
type Product struct {
	ID             string    `json:"id"`
	ManufacturerID string    `json:"manufacturerId"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]Product, error)
	GetProductByID(ctx context.Context, id string) (*Product, error)
	CreateProduct(ctx context.Context, product Product) (*Product, error)
	UpdateProduct(ctx context.Context, product Product) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

// ProductServiceImpl - реализация сервиса для работы с продуктами.
type ProductServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

// NewProductService - функция для создания нового сервиса продуктов.
func NewProductService(db *sql.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		db: db,
	}
}

// GetAllProducts - получение списка всех продуктов.
func (s *ProductServiceImpl) GetAllProducts(ctx context.Context) ([]Product, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.ManufacturerID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductByID - получение продукта по его ID.
func (s *ProductServiceImpl) GetProductByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	err := s.db.QueryRowContext(ctx, "SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.ManufacturerID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("продукт не найден")
		}
		return nil, err
	}
	return &product, nil
}

// CreateProduct - создание нового продукта.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO products (manufacturerId, name, description, price) VALUES ($1, $2, $3, $4)", product.ManufacturerID, product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = fmt.Sprintf("%d", lastInsertID)
	return &product, nil
}

// UpdateProduct - обновление продукта.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, product Product) (*Product, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE products SET manufacturerId = $1, name = $2, description = $3, price = $4 WHERE id = $5", product.ManufacturerID, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	return &product, nil // Возвращаем обновленный продукт
}

// DeleteProduct - удаление продукта.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
