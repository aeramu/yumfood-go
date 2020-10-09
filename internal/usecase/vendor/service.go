package vendor

import (
	"log"

	"github.com/aeramu/yumfood-go/internal/domain"
)

type service struct {
	repo  Repository
	idGen IDGenerator
}

func (s service) CreateVendor(name, description string) {
	err := s.repo.PutVendor(name, description)
	if err != nil {
		log.Println(err)
	}
}

func (s service) UpdateVendor(id, name, description string) {
	err := s.repo.UpdateVendor(id, name, description)
	if err != nil {
		log.Println(err)
	}
}

func (s service) GetVendor(id string) domain.Vendor {
	vendor, err := s.repo.GetVendor(id)
	if err != nil {
		log.Println(err)
	}
	return vendor
}

func (s service) DeleteVendor(id string) {
	err := s.repo.DeleteVendor(id)
	if err != nil {
		log.Println(err)
	}
}
