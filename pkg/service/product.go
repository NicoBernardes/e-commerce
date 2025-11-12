package service

import (
	"e-commerce/pkg/handlers"
	"time"
)

type ProductService struct {
	ProductID   string  `validade:"required" gorm:"primaryKey;size:50;not null"`
	ProductName string  `validade:"required" gorm:"size:100;not null"`
	Price       float64 `validade:"required" gorm:"not null"`
	Stock       int     `validade:"required" gorm:"not null"`
	Category    string  `validade:"required" gorm:"size:30"`
	Sale        bool
	UpdatedAt   time.Time
}

func NewProduct(name string, price float64, stock int, category string, sale bool) (*ProductService, error) {

	product := &ProductService{
		ProductName: name,
		Price:       price,
		Stock:       stock,
		Category:    category,
		Sale:        sale,
		UpdatedAt:   time.Now(),
	}

	err := handlers.ValidateStruct(product)
	if err == nil {
		return product, nil
	}

	return nil, err
}
