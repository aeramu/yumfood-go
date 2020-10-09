package order

import "github.com/aeramu/yumfood-go/internal/domain"

// Service interface as a port
type Service interface {
	Create(items []domain.OrderItem)
	GetList() []domain.Order
}

func NewService(repo Repository, idGen IDGenerator) Service {
	return &service{
		repo:  repo,
		idGen: idGen,
	}
}
