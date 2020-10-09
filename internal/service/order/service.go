package order

import (
	"log"

	"github.com/aeramu/yumfood-go/internal/domain"
)

type service struct {
	repo  Repository
	idGen IDGenerator
}

func (s service) Create(items []domain.OrderItem) {
	id := s.idGen.Generate()
	err := s.repo.Put(id, items)
	if err != nil {
		log.Println(err)
	}
}

func (s service) GetList() []domain.Order {
	orders, err := s.repo.GetAll()
	if err != nil {
		log.Println(err)
	}
	return orders
}
