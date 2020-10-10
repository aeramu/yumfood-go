package vendor

import (
	"github.com/aeramu/yumfood-go/internal/domain"
	"github.com/aeramu/yumfood-go/internal/service/order"
)

// NewRepository initiate repository implementation
func NewRepository() order.Repository {
	orders := []domain.Order{
		{"1", []domain.OrderItem{
			{"1", 18, "sangat pedas"},
			{"2", 2, "jangan pedas"},
		}},
		{"2", []domain.OrderItem{
			{"4", 18, "sangat pedas"},
			{"5", 18, "sangat pedas"},
			{"6", 1, "sangat sangat pedas"},
		}},
		{"3", []domain.OrderItem{
			{"4", 18, "sangat pedas"},
		}},
	}
	return &repository{
		data: orders,
	}
}

type repository struct {
	data []domain.Order
}

func (r *repository) Put(id string, items []domain.OrderItem) error {
	r.data = append(r.data, domain.Order{id, items})
	return nil
}

func (r *repository) GetAll() ([]domain.Order, error) {
	return r.data, nil
}
