package repository

import (
	"bagstore/internal/models"
	"gorm.io/gorm"
)

type BagRepositoryImpl struct {
	db *gorm.DB
}

// NewBagRepository - Конструктор
func NewBagRepository(db *gorm.DB) *BagRepositoryImpl {
	return &BagRepositoryImpl{db: db}
}

// GetAll - Получить все сумки
func (r BagRepositoryImpl) GetAll() ([]models.Bag, error) {
	var bags []models.Bag
	err := r.db.Find(&bags).Error
	return bags, err
}

// GetById - Получить сумку по ID
func (r BagRepositoryImpl) GetById(id int) (*models.Bag, error) {
	var bag models.Bag
	err := r.db.First(&bag, id).Error
	return &bag, err
}

// Create - Добавить новую сумку
func (r BagRepositoryImpl) Create(bag *models.Bag) error {
	return r.db.Create(bag).Error
}

// Update - Обновить данные сумки
func (r BagRepositoryImpl) Update(id int, bagEdit *models.BagEdit) error {
	// Ожидается, что BagEdit будет использоваться для обновления данных
	// Мы исключаем поля "id" и "CreatedAt", чтобы их не изменять
	return r.db.Model(&models.Bag{}).Where("id = ?", id).Omit("id", "CreatedAt").Updates(bagEdit).Error
}

// Delete - Удалить сумку по ID
func (r BagRepositoryImpl) Delete(bagID int) error {
	var bag models.Bag
	// Ищем сумку по ID перед удалением
	if err := r.db.First(&bag, bagID).Error; err != nil {
		return err // Если сумка не найдена, возвращаем ошибку
	}

	// Удаляем найденную сумку
	return r.db.Delete(&bag).Error
}
