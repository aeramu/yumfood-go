package vendor

import (
	"log"

	"github.com/aeramu/yumfood-go/internal/domain"
)

type service struct {
	repo  Repository
	idGen IDGenerator
}

func (s service) Create(name, description string) {
	id := s.idGen.Generate()
	if len(name) >= 128 {
		name = name[:127]
	}
	err := s.repo.Put(id, name, description)
	if err != nil {
		log.Println(err)
	}
}

func (s service) Update(id, name, description string) {
	if len(name) >= 128 {
		name = name[:127]
	}
	err := s.repo.Update(id, name, description)
	if err != nil {
		log.Println(err)
	}
}

func (s service) Get(id string) domain.Vendor {
	vendor, err := s.repo.Get(id)
	if err != nil {
		log.Println(err)
	}
	return vendor
}

func (s service) Delete(id string) {
	err := s.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
}
