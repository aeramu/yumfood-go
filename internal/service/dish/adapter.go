package dish

import "github.com/aeramu/yumfood-go/internal/domain"

// Repository interface as adapter
type Repository interface {
	Put(id, vendorID, name string, price int) error
	Update(id, name string, price int) error
	Get(id string) (domain.Dish, error)
	GetListByVendorID(vendorID string) ([]domain.Dish, error)
	Delete(id string) error
}

// IDGenerator interface as adapter
type IDGenerator interface {
	Generate() string
}
