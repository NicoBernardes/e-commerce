package service

type Repository interface {
	CreateProduct(product *ProductService) error
	GetAllProducts() ([]ProductService, error)
	GetProductByID(id uint) (*ProductService, error)
	UpdateProduct(product *ProductService) error
	DeleteProduct(id uint) error
}
