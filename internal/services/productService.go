package services //заменить s.db на фактическое подключение к бд, но это в конце после маина

import (
	"context"
	"errors"
	"time"

	"database/sql"
)

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
	GetAllProducts() ([]Product, error)
	GetProductByID(id string) (*Product, error)
	CreateProduct(product Product) (*Product, error)
	UpdateProduct(product Product) error
	DeleteProduct(id string) error
}

type ProductServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewProductService(db *sql.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		db: db,
	}
}

func (s *ProductServiceImpl) GetAllProducts() ([]Product, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM products")
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

func (s *ProductServiceImpl) GetProductByID(id string) (*Product, error) {
	var product Product
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.ManufacturerID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("продукт не найден")
		}
		return nil, err
	}
	return &product, nil
}

func (s *ProductServiceImpl) CreateProduct(product Product) (*Product, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO products (manufacturerId, name, description, price) VALUES ($1, $2, $3, $4)", product.ManufacturerID, product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	product.ID = lastInsertID
	return &product, nil
}

func (s *ProductServiceImpl) UpdateProduct(product Product) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE products SET manufacturerId = $1, name = $2, description = $3, price = $4 WHERE id = $5", product.ManufacturerID, product.Name, product.Description, product.Price, product.ID)
	return err
}

func (s *ProductServiceImpl) DeleteProduct(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
