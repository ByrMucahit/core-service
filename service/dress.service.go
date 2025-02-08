package service

import (
	"core-service/models"
	"core-service/repository"
)

type DressService struct {
	repo *repository.DressRepository
}

func DressServiceInstance() *DressService {
	dressRepo := repository.NewProductRepository(repository.DB)

	return &DressService{
		repo: dressRepo,
	}
}

func (s *DressService) AddDress(body models.Dress) error {
	return s.repo.CreateDress(&body)
}
