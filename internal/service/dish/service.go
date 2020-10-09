package dish

import (
	"log"

	"github.com/aeramu/yumfood-go/internal/domain"
)

type service struct {
	repo  Repository
	idGen IDGenerator
}

func (s service) Create(vendorID, name string, price int) {
	id := s.idGen.Generate()
	err := s.repo.Put(id, vendorID, name, price)
	if err != nil {
		log.Println(err)
	}
}

func (s service) Update(id, name string, price int) {
	err := s.repo.Update(id, name, price)
	if err != nil {
		log.Println(err)
	}
}

func (s service) Get(id string) domain.Dish {
	dish, err := s.repo.Get(id)
	if err != nil {
		log.Println(err)
	}
	return dish
}

func (s service) Delete(id string) {
	err := s.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
}

func (s service) GetListByVendorID(vendorID string) []domain.Dish {
	dishes, err := s.repo.GetListByVendorID(vendorID)
	if err != nil {
		log.Println(err)
	}
	return dishes
}
