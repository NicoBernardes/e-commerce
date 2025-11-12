package service

type ProductService struct {
	ProductID   string  `validade:"required" gorm:"primaryKey;size:50;not null"`
	ProductName string  `validade:"required" gorm:"size:100;not null"`
	Price       float64 `validade:"required" gorm:"not null"`
	Stock       int     `validade:"required" gorm:"not null"`
	Sale        bool
	Category    string `validade:"required" gorm:"size:30"`
}
