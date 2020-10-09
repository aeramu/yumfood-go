package order

import "github.com/aeramu/yumfood-go/internal/domain"

// Repository interface as adapter
type Repository interface {
	Put(id string, items []domain.OrderItem) error
	GetAll() ([]domain.Order, error)
}

// IDGenerator interface as adapter
type IDGenerator interface {
	Generate() string
}
