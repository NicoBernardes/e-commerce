package service

type OrderRepository interface {
	CreateOrder(order *OrderService) error
	GetAllOrders() ([]OrderService, error)
	GetOrderByID(id uint) (*OrderService, error)
	UpdateOrder(order *OrderService) error
	DeleteOrder(id uint) error
}
