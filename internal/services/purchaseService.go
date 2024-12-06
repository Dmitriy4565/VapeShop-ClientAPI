package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type Purchase struct {
	ID         int64 `json:"id"`
	CustomerID int64 `json:"customer_id"`
	StoreID    int64 `json:"store_id"`
	ProductID  int64 `json:"product_id"`
	Quantity   int64 `json:"quantity"` // Изменено на NullInt64
}

type PurchaseService interface {
	GetAllPurchases(ctx context.Context) ([]Purchase, error)
	GetPurchaseByID(ctx context.Context, id string) (*Purchase, error)
	CreatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error)
	UpdatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error)
	DeletePurchase(ctx context.Context, id string) error
}

type PurchaseServiceImpl struct {
	db *sql.DB
}

func NewPurchaseService(db *sql.DB) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		db: db,
	}
}

func (s *PurchaseServiceImpl) GetAllPurchases(ctx context.Context) ([]Purchase, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, customer_id, store_id, product_id, quantity FROM purchases")
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer rows.Close()

	var purchases []Purchase
	for rows.Next() {
		var purchase Purchase
		if err := rows.Scan(
			&purchase.ID,
			&purchase.CustomerID,
			&purchase.StoreID,
			&purchase.ProductID,
			&purchase.Quantity,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}

		purchases = append(purchases, purchase)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам: %w", err)
	}

	return purchases, nil
}

func (s *PurchaseServiceImpl) GetPurchaseByID(ctx context.Context, id string) (*Purchase, error) {
	var purchase Purchase

	// Преобразование ID из строки в целое число
	purchaseID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("недопустимый ID")
	}

	// Выполнение запроса к базе данных с явным указанием столбцов
	err = s.db.QueryRowContext(ctx, "SELECT id, customer_id, store_id, product_id, quantity FROM purchases WHERE id = $1", purchaseID).Scan(
		&purchase.ID,
		&purchase.CustomerID,
		&purchase.StoreID,
		&purchase.ProductID,
		&purchase.Quantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("покупка не найдена")
		}
		return nil, err // Возврат ошибки при других проблемах
	}

	return &purchase, nil // Возврат найденной покупки
}

func (s *PurchaseServiceImpl) CreatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error) {
	// Используем RETURNING для получения ID вставленной записи
	query := `
        INSERT INTO purchases (customer_id, store_id, product_id, quantity) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`

	err := s.db.QueryRowContext(ctx, query,
		purchase.CustomerID,
		purchase.StoreID,
		purchase.ProductID,
		purchase.Quantity).Scan(&purchase.ID)

	if err != nil {
		return nil, err // Возврат ошибки при выполнении запроса
	}

	return &purchase, nil // Возврат созданной покупки с установленным ID
}

func (s *PurchaseServiceImpl) UpdatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error) {
	res, err := s.db.ExecContext(ctx, `
        UPDATE purchases 
        SET customer_id = $1, store_id = $2, product_id = $3, quantity = $4 
        WHERE id = $5`,
		purchase.CustomerID, purchase.StoreID, purchase.ProductID, purchase.Quantity, purchase.ID)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("запись не найдена") // Обработка случая, когда запись не найдена
	}

	return &purchase, nil // Возвращаем обновлённую покупку
}

func (s *PurchaseServiceImpl) DeletePurchase(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM purchases WHERE id = $1", id)
	return err
}
