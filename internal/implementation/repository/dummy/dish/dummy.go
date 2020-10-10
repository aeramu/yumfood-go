package vendor

import (
	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/dish"
)

// NewRepository initiate repository implementation
func NewRepository() dish.Repository {
	dishes := []domain.Dish{
		{"1", "1", "old cuisine", 3000},
		{"2", "1", "new cuisine", 7000},
		{"3", "2", "pizza mama", 23000},
		{"4", "3", "crabs", 5000},
		{"5", "3", "crabs", 53000},
		{"6", "3", "crabs", 15000},
	}
	return &repository{
		data: dishes,
	}
}

type repository struct {
	data []domain.Dish
}

func (r *repository) Put(id, vendorID, name string, price int) error {
	r.data = append(r.data, domain.Dish{id, vendorID, name, price})
	return nil
}

func (r *repository) Update(id, name string, price int) error {
	for i, dish := range r.data {
		if dish.ID == id {
			if name != "" {
				r.data[i].Name = name
			}
			if price != -1 {
				r.data[i].Price = price
			}
		}
	}
	return nil
}
func (r *repository) Get(id string) (domain.Dish, error) {
	for _, dish := range r.data {
		if dish.ID == id {
			return dish, nil
		}
	}
	return domain.Dish{}, nil
}

func (r *repository) Delete(id string) error {
	for i, dish := range r.data {
		if dish.ID == id {
			copy(r.data[i:], r.data[i+1:])
			r.data = r.data[:len(r.data)-1]
		}
	}
	return nil
}

func (r *repository) GetListByVendorID(vendorID string) ([]domain.Dish, error) {
	var dishes []domain.Dish
	for _, dish := range r.data {
		if dish.VendorID == vendorID {
			dishes = append(dishes, dish)
		}
	}
	return dishes, nil
}
