package vendor

import "github.com/aeramu/yumfood-go/internal/domain"

// Service port for vendor service
type Service interface {
	CreateVendor(name, description string)
	UpdateVendor(id, name, description string)
	GetVendor(id string) domain.Vendor
	DeleteVendor(id string)
}

// NewService constructor for service
func NewService(repo Repository, idGen IDGenerator) Service {
	return &service{
		repo:  repo,
		idGen: idGen,
	}
}
