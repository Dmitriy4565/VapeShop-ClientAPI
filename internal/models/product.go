package models

// Product представляет собой структуру продукта
type Product struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	ImageURL        string  `json:"image_url"`
	CategoryID      int     `json:"category_id"`
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

// NewProduct создает новый продукт
func NewProduct(name, description string, price float64, imageURL string, categoryID, manufacturerID, stock int, vapeType string, power, batteryCapacity int, tankCapacity, coilResistance float64, material, color string, isNew, isFeatured bool) *Product {
	return &Product{
		Name:            name,
		Description:     description,
		Price:           price,
		ImageURL:        imageURL,
		CategoryID:      categoryID,
		ManufacturerID:  manufacturerID,
		Stock:           stock,
		VapeType:        vapeType,
		Power:           power,
		BatteryCapacity: batteryCapacity,
		TankCapacity:    tankCapacity,
		CoilResistance:  coilResistance,
		Material:        material,
		Color:           color,
		IsNew:           isNew,
		IsFeatured:      isFeatured,
	}
}

// Update обновляет данные продукта
func (p *Product) Update(name, description string, price float64, imageURL string, categoryID, manufacturerID, stock int, vapeType string, power, batteryCapacity int, tankCapacity, coilResistance float64, material, color string, isNew, isFeatured bool) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.ImageURL = imageURL
	p.CategoryID = categoryID
	p.ManufacturerID = manufacturerID
	p.Stock = stock
	p.VapeType = vapeType
	p.Power = power
	p.BatteryCapacity = batteryCapacity
	p.TankCapacity = tankCapacity
	p.CoilResistance = coilResistance
	p.Material = material
	p.Color = color
	p.IsNew = isNew
	p.IsFeatured = isFeatured
}
