package repository

import (
	"core-service/models"
	"gorm.io/gorm"
)

type DressRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *DressRepository {
	return &DressRepository{DB: db}
}

func (r *DressRepository) CreateDress(dress *models.Dress) error {
	return r.DB.Create(dress).Error
}
