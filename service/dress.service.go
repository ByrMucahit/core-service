package service

import (
	"core-service/models"
	"core-service/repository"
)

type DressService struct {
	DressRepo *repository.DressRepository
}

func DressServiceInstance(repo *repository.DressRepository) *DressService {
	return &DressService{DressRepo: repo}
}

func (s *DressService) AddDress(body models.Dress) error {
	return s.DressRepo.CreateDress(&body)
}

func (s *DressService) FindDress() ([]models.Dress, error) {
	return s.DressRepo.FindDress()
}

func (s *DressService) UpdatesPartial(entity *models.Dress, id uint, updates map[string]interface{}) error {
	return s.DressRepo.UpdatePartialDress(entity, id, updates)
}

func (s *DressService) DeletePartialDress(entity *models.Dress, id uint) error {
	return s.DressRepo.DeletePartialDress(entity, id)
}
