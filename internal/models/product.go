package models

import (
	"time"
)

type Product struct {
	ID             int       `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Price          float64   `json:"price" db:"price"`
	ImageURL       string    `json:"image_url" db:"image_url"`
	CategoryID     int       `json:"category_id" db:"category_id"`
	ManufacturerID int       `json:"manufacturer_id" db:"manufacturer_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(name, description string, price float64, imageURL string, categoryID, manufacturerID int) *Product {
	return &Product{
		Name:           name,
		Description:    description,
		Price:          price,
		ImageURL:       imageURL,
		CategoryID:     categoryID,
		ManufacturerID: manufacturerID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (p *Product) Update(name, description string, price float64, imageURL string, categoryID, manufacturerID int) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.ImageURL = imageURL
	p.CategoryID = categoryID
	p.ManufacturerID = manufacturerID
	p.UpdatedAt = time.Now()
}
