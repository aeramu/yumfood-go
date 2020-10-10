package vendorrepo

import (
	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/vendor"
)

// NewRepository initiate repository implementation
func NewRepository() vendor.Repository {
	vendors := []domain.Vendor{
		{"1", "basingse", "old cuisine"},
		{"2", "mama pizza", "hot chili sauce with pizza mama"},
		{"3", "rest-rest-resto", "restaurant with sea aroma"},
	}
	return &repository{
		data: vendors,
	}
}

type repository struct {
	data []domain.Vendor
}

func (r *repository) Put(id, name, description string) error {
	r.data = append(r.data, domain.Vendor{id, name, description})
	return nil
}

func (r *repository) Update(id, name, description string) error {
	for i, vendor := range r.data {
		if vendor.ID == id {
			if name != "" {
				r.data[i].Name = name
			}
			if description != "" {
				r.data[i].Description = description
			}
		}
	}
	return nil
}
func (r *repository) Get(id string) (domain.Vendor, error) {
	for _, vendor := range r.data {
		if vendor.ID == id {
			return vendor, nil
		}
	}
	return domain.Vendor{}, nil
}

func (r *repository) Delete(id string) error {
	for i, vendor := range r.data {
		if vendor.ID == id {
			copy(r.data[i:], r.data[i+1:])
			r.data = r.data[:len(r.data)-1]
		}
	}
	return nil
}
