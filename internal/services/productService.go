package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// Product - структура, представляющая продукт.
type Product struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	ImageUrl        string  `json:"image_url"`
	CategoryID      int     `json:"category_id"` // Убедитесь, что это поле правильно указано
	ManufacturerID  int     `json:"manufacturer_id"`
	Stock           int     `json:"stock"`
	VapeType        string  `json:"vape_type"`
	Power           int     `json:"power"`
	BatteryCapacity int     `json:"battery_capacity"`
	TankCapacity    float64 `json:"tank_capacity"`
	CoilResistance  float64 `json:"coil_resistance"`
	Material        string  `json:"material"`
	Color           string  `json:"color"`
	IsNew           bool    `json:"is_new"`
	IsFeatured      bool    `json:"is_featured"`
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
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageUrl,
			&product.CategoryID,
			&product.ManufacturerID,
			&product.Stock,
			&product.VapeType,
			&product.Power,
			&product.BatteryCapacity,
			&product.TankCapacity,
			&product.CoilResistance,
			&product.Material,
			&product.Color,
			&product.IsNew,
			&product.IsFeatured,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %w", err)
	}

	return products, nil
}

// GetProductByID - получение продукта по его ID.
func (s *ProductServiceImpl) GetProductByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	err := s.db.QueryRowContext(ctx, "SELECT * FROM products WHERE id = $1", id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageUrl,
		&product.CategoryID,
		&product.ManufacturerID,
		&product.Stock,
		&product.VapeType,
		&product.Power,
		&product.BatteryCapacity,
		&product.TankCapacity,
		&product.CoilResistance,
		&product.Material,
		&product.Color,
		&product.IsNew,
		&product.IsFeatured,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("продукт не найден")
		}
		return nil, fmt.Errorf("ошибка получения продукта по ID: %w", err)
	}
	return &product, nil
}

// CreateProduct - создание нового продукта.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	// Проверка существования категории
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM categories WHERE id=$1)", product.CategoryID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки категории: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("категория с ID %d не существует", product.CategoryID)
	}

	// Подготовка SQL-запроса
	query := `
        INSERT INTO products (name, description, price, image_url, category_id, manufacturer_id, stock, vape_type, power, battery_capacity, tank_capacity, coil_resistance, material, color, is_new, is_featured)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
        RETURNING id`

	// Выполнение запроса
	res := s.db.QueryRowContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.ImageUrl,
		product.CategoryID,
		product.ManufacturerID,
		product.Stock,
		product.VapeType,
		product.Power,
		product.BatteryCapacity,
		product.TankCapacity,
		product.CoilResistance,
		product.Material,
		product.Color,
		product.IsNew,
		product.IsFeatured)

	// Получение ID последней вставленной записи
	if err := res.Scan(&product.ID); err != nil {
		return nil, fmt.Errorf("ошибка при вставке продукта: %w", err)
	}

	return &product, nil
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, product Product) (*Product, error) {
	fmt.Printf("Обновление продукта с ID: %d\n", product.ID) // Логируем ID

	result, err := s.db.ExecContext(ctx, `
        UPDATE products 
        SET manufacturer_id = $1, name = $2, description = $3, price = $4 
        WHERE id = $5`,
		product.ManufacturerID, product.Name, product.Description, product.Price, product.ID)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("продукт с ID %d не найден", product.ID)
	}

	return &product, nil // Возвращаем обновленный продукт
}

// DeleteProduct - удаление продукта.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
