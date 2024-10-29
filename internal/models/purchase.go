package models

import (
	"time"
)

type Purchase struct {
	ID         int       `json:"id" db:"id"`
	CustomerID int       `json:"customer_id" db:"customer_id"`
	ProductID  int       `json:"product_id" db:"product_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	TotalPrice float64   `json:"total_price" db:"total_price"`
	DeliveryID int       `json:"delivery_id" db:"delivery_id"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func NewPurchase(customerID, productID, quantity int, totalPrice float64, deliveryID, status string) *Purchase {
	return &Purchase{
		CustomerID: customerID,
		ProductID:  productID,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		DeliveryID: deliveryID,
		Status:     status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (p *Purchase) Update(customerID, productID, quantity int, totalPrice float64, deliveryID, status string) {
	p.CustomerID = customerID
	p.ProductID = productID
	p.Quantity = quantity
	p.TotalPrice = totalPrice
	p.DeliveryID = deliveryID
	p.Status = status
	p.UpdatedAt = time.Now()
}
