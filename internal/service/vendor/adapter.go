package vendor

import "github.com/aeramu/yumfood-go/internal/domain"

// Repository interface as adapter
type Repository interface {
	Put(id, name, description string) error
	Update(id, name, description string) error
	Get(id string) (domain.Vendor, error)
	GetAll() ([]domain.Vendor, error)
	GetListByTags(tags []string) ([]domain.Vendor, error)
	Delete(id string) error
	Tag(id, tag string) error
}

// IDGenerator interface as adapter
type IDGenerator interface {
	Generate() string
}
