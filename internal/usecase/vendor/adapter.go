package vendor

import "github.com/aeramu/yumfood-go/internal/domain"

// Repository interface as adapter
type Repository interface {
	PutVendor(name, description string) error
	UpdateVendor(id, name, description string) error
	GetVendor(id string) (domain.Vendor, error)
	DeleteVendor(id string) error
}

// IDGenerator interface as adapter
type IDGenerator interface {
	Generate() string
}
