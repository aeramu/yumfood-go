package vendor

import (
	"log"

	"github.com/aeramu/yumfood-go/internal/domain"
)

type service struct {
	repo  Repository
	idGen IDGenerator
}

func (s service) Create(name, description string, tags ...string) {
	id := s.idGen.Generate()
	if len(name) >= 128 {
		name = name[:127]
	}
	err := s.repo.Put(id, name, description)
	if err != nil {
		log.Println(err)
	}
	for _, tag := range tags {
		err = s.repo.Tag(id, tag)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s service) Update(id, name, description string, tags ...string) {
	if len(name) >= 128 {
		name = name[:127]
	}
	err := s.repo.Update(id, name, description)
	if err != nil {
		log.Println(err)
	}
	for _, tag := range tags {
		err = s.repo.Tag(id, tag)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s service) Get(id string) domain.Vendor {
	vendor, err := s.repo.Get(id)
	if err != nil {
		log.Println(err)
	}
	return vendor
}

func (s service) GetList(tags ...string) []domain.Vendor {
	var vendors []domain.Vendor
	var err error
	if len(tags) > 0 {
		vendors, err = s.repo.GetListByTags(tags)
	} else {
		vendors, err = s.repo.GetAll()
	}

	if err != nil {
		log.Println(err)
	}
	return vendors
}

func (s service) Delete(id string) {
	err := s.repo.Delete(id)
	if err != nil {
		log.Println(err)
	}
}
