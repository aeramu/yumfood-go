package dish

import "github.com/aeramu/yumfood-go/internal/domain"

// Service port for vendor service
type Service interface {
	Create(vendorID, name string, price int)
	Get(id string) domain.Dish
	Update(id, name string, price int)
	Delete(id string)
	GetListByVendorID(vendorID string) []domain.Dish
}

// NewService constructor for service
func NewService(repo Repository, idGen IDGenerator) Service {
	return &service{
		repo:  repo,
		idGen: idGen,
	}
}
