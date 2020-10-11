package vendor

import "github.com/aeramu/yumfood-go/internal/domain"

// Service port for vendor service
type Service interface {
	Create(name, description string, tags ...string)
	Update(id, name, description string, tags ...string)
	Get(id string) domain.Vendor
	GetList(tags ...string) []domain.Vendor
	Delete(id string)
}

// NewService constructor for service
func NewService(repo Repository, idGen IDGenerator) Service {
	return &service{
		repo:  repo,
		idGen: idGen,
	}
}
