package services //заменить s.db на фактическое подключение к бд, но это в конце после маина

import (
	"context"
	"errors"
	"time"

	"database/sql"
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
	GetAllPurchases() ([]Purchase, error)
	GetPurchaseByID(id string) (*Purchase, error)
	CreatePurchase(purchase Purchase) (*Purchase, error)
	UpdatePurchase(purchase Purchase) error
	DeletePurchase(id string) error
}

type PurchaseServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewPurchaseService(db *sql.DB) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		db: db,
	}
}

func (s *PurchaseServiceImpl) GetAllPurchases() ([]Purchase, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM purchases")
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

func (s *PurchaseServiceImpl) GetPurchaseByID(id string) (*Purchase, error) {
	var purchase Purchase
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM purchases WHERE id = $1", id).Scan(&purchase.ID, &purchase.CustomerID, &purchase.StoreID, &purchase.ProductID, &purchase.Quantity, &purchase.CreatedAt, &purchase.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("покупка не найдена")
		}
		return nil, err
	}
	return &purchase, nil
}

func (s *PurchaseServiceImpl) CreatePurchase(purchase Purchase) (*Purchase, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO purchases (customerId, storeId, productId, quantity) VALUES ($1, $2, $3, $4)", purchase.CustomerID, purchase.StoreID, purchase.ProductID, purchase.Quantity)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	purchase.ID = lastInsertID
	return &purchase, nil
}

func (s *PurchaseServiceImpl) UpdatePurchase(purchase Purchase) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE purchases SET customerId = $1, storeId = $2, productId = $3, quantity = $4 WHERE id = $5", purchase.CustomerID, purchase.StoreID, purchase.ProductID, purchase.Quantity, purchase.ID)
	return err
}

func (s *PurchaseServiceImpl) DeletePurchase(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM purchases WHERE id = $1", id)
	return err
}
