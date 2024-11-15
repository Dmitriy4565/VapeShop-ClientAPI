package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Purchase struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerId"`
	StoreID    string    `json:"storeId"`
	ProductID  string    `json:"productId"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
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
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM purchases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []Purchase
	for rows.Next() {
		var purchase Purchase
		if err := rows.Scan(&purchase.ID, &purchase.CustomerID, &purchase.StoreID, &purchase.ProductID, &purchase.Quantity, &purchase.CreatedAt, &purchase.UpdatedAt); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}

func (s *PurchaseServiceImpl) GetPurchaseByID(ctx context.Context, id string) (*Purchase, error) {
	var purchase Purchase
	err := s.db.QueryRowContext(ctx, "SELECT * FROM purchases WHERE id = $1", id).Scan(&purchase.ID, &purchase.CustomerID, &purchase.StoreID, &purchase.ProductID, &purchase.Quantity, &purchase.CreatedAt, &purchase.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("покупка не найдена")
		}
		return nil, err
	}
	return &purchase, nil
}

func (s *PurchaseServiceImpl) CreatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO purchases (customerId, storeId, productId, quantity) VALUES ($1, $2, $3, $4)", purchase.CustomerID, purchase.StoreID, purchase.ProductID, purchase.Quantity)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	purchase.ID = fmt.Sprintf("%d", lastInsertID)
	return &purchase, nil
}

func (s *PurchaseServiceImpl) UpdatePurchase(ctx context.Context, purchase Purchase) (*Purchase, error) {
	res, err := s.db.ExecContext(ctx, "UPDATE purchases SET customerId = $1, storeId = $2, productId = $3, quantity = $4 WHERE id = $5", purchase.CustomerID, purchase.StoreID, purchase.ProductID, purchase.Quantity, purchase.ID)
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
