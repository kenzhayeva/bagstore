package repository

import (
	"bagstore/internal/models"
	"bagstore/internal/services"
	"errors"
	"gorm.io/gorm"
)

type BagRepositoryImpl struct {
	db *gorm.DB
}

func NewBagRepository(db *gorm.DB) *BagRepositoryImpl {
	return &BagRepositoryImpl{db: db}
}

func (r *BagRepositoryImpl) GetAll(userID uint) ([]models.Bag, error) {
	var bags []models.Bag
	if err := r.db.Where("user_id = ?", userID).Find(&bags).Error; err != nil {
		return nil, err
	}
	return bags, nil
}

func (r *BagRepositoryImpl) GetByID(userID, bagID uint) (*models.Bag, error) {
	var bag models.Bag
	err := r.db.Where("user_id = ? AND id = ?", userID, bagID).First(&bag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, services.ErrBagNotFound
	}
	return &bag, err
}

func (r *BagRepositoryImpl) Create(bag *models.Bag) error {
	if err := r.db.Create(bag).Error; err != nil {
		return err
	}
	return nil
}

func (r *BagRepositoryImpl) Update(userID, bagID uint, bagEdit *models.BagEdit) error {
	result := r.db.Model(&models.Bag{}).
		Where("user_id = ? AND id = ?", userID, bagID).
		Updates(models.Bag{
			Title:    bagEdit.Title,
			Category: bagEdit.Category,
			Color:    bagEdit.Color,
			Price:    bagEdit.Price,
			Size:     bagEdit.Size,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return services.ErrBagNotFound
	}
	return nil
}

func (r *BagRepositoryImpl) Delete(userID, bagID uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, bagID).Delete(&models.Bag{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return services.ErrBagNotFound
	}
	return nil
}
