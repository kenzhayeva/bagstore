<<<<<<< HEAD
=======
//package repository
//
//import (
//	"bagstore/internal/models"
//	"gorm.io/gorm"
//)
//
//type BagRepositoryImpl struct {
//	db *gorm.DB
//}
//
//func NewBagRepository(db *gorm.DB) *BagRepositoryImpl {
//	return &BagRepositoryImpl{db: db}
//}
//
//func (r BagRepositoryImpl) GetAll() ([]models.Bag, error) {
//	var bags []models.Bag
//	err := r.db.Find(&bags).Error
//	return bags, err
//}
//
//func (r BagRepositoryImpl) GetById(id int) (*models.Bag, error) {
//	var bag models.Bag
//	err := r.db.First(&bag, id).Error
//	return &bag, err
//}
//
//func (r BagRepositoryImpl) Create(bag *models.Bag) error {
//	return r.db.Create(bag).Error
//}
//
//func (r BagRepositoryImpl) Update(id int, bagEdit *models.BagEdit) error {
//	return r.db.Model(&models.Bag{}).Where("id = ?", id).Omit("id", "CreatedAt").Updates(bagEdit).Error
//}
//
//func (r BagRepositoryImpl) Delete(bagID int) error {
//	var bag models.Bag
//	// Ищем сумку по ID перед удалением
//	if err := r.db.First(&bag, bagID).Error; err != nil {
//		return err // Если сумка не найдена, возвращаем ошибку
//	}
//
//	return r.db.Delete(&bag).Error
//}

>>>>>>> be3f14a (Initial commit)
package repository

import (
	"bagstore/internal/models"
	"gorm.io/gorm"
)

type BagRepositoryImpl struct {
	db *gorm.DB
}

<<<<<<< HEAD
// NewBagRepository - Конструктор
=======
>>>>>>> be3f14a (Initial commit)
func NewBagRepository(db *gorm.DB) *BagRepositoryImpl {
	return &BagRepositoryImpl{db: db}
}

<<<<<<< HEAD
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
=======
// Получить все сумки конкретного пользователя
func (r BagRepositoryImpl) GetAllByUser(userID uint) ([]models.Bag, error) {
	var bags []models.Bag
	err := r.db.Where("user_id = ?", userID).Find(&bags).Error
	return bags, err
}

// Получить сумку по ID, но только если она принадлежит пользователю
func (r BagRepositoryImpl) GetById(userID, bagID uint) (*models.Bag, error) {
	var bag models.Bag
	err := r.db.Where("id = ? AND user_id = ?", bagID, userID).First(&bag).Error
	return &bag, err
}

// Создать сумку (UserID должен быть установлен в handler-е)
>>>>>>> be3f14a (Initial commit)
func (r BagRepositoryImpl) Create(bag *models.Bag) error {
	return r.db.Create(bag).Error
}

<<<<<<< HEAD
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
=======
// Обновить сумку только если она принадлежит пользователю
func (r BagRepositoryImpl) Update(userID, bagID uint, bagEdit *models.BagEdit) error {
	return r.db.Model(&models.Bag{}).
		Where("id = ? AND user_id = ?", bagID, userID).
		Omit("id", "CreatedAt").
		Updates(bagEdit).Error
}

// Удалить сумку, только если она принадлежит пользователю
func (r BagRepositoryImpl) Delete(userID, bagID uint) error {
	return r.db.Where("id = ? AND user_id = ?", bagID, userID).
		Delete(&models.Bag{}).Error
>>>>>>> be3f14a (Initial commit)
}
