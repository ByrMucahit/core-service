package repository

import (
	"core-service/models"
	"errors"
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

func (r *DressRepository) FindDress() ([]models.Dress, error) {
	var dresses []models.Dress
	err := r.DB.Find(&dresses).Error
	if err != nil {
		return nil, err
	}
	return dresses, nil
}

func (r *DressRepository) UpdatePartialDress(entity *models.Dress, id uint, updates map[string]interface{}) error {
	err := r.DB.Model(entity).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DressRepository) DeletePartialDress(entity *models.Dress, id uint) error {
	result := r.DB.Debug().Where("id = ? AND deleted_at IS NULL", id).Delete(entity)
	if result.RowsAffected == 0 {
		return errors.New("record not found or already deleted")
	}
	return result.Error
}
