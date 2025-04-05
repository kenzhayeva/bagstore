package services

import (
	"bagstore/internal/models"
)

// Интерфейс репозитория сумок
type BagRepository interface {
	GetAll() ([]models.Bag, error)
	GetById(id int) (*models.Bag, error)
	Create(bag *models.Bag) error
	Update(id int, bagEdit *models.BagEdit) error
	Delete(bagID int) error
}

// Структура сервиса сумок
type BagService struct {
	repo BagRepository
}

// Конструктор BagService
func NewBagService(bagRepo BagRepository) *BagService {
	return &BagService{repo: bagRepo}
}

// Получение всех сумок
func (s *BagService) GetAllBags() ([]models.Bag, error) {
	return s.repo.GetAll()
}

// Получение сумки по ID
func (s *BagService) GetBagByID(id int) (*models.Bag, error) {
	return s.repo.GetById(id)
}

// Создание новой сумки
func (s *BagService) Create(bagEdit models.BagEdit) (*models.Bag, error) {
	bag := &models.Bag{
		Title:    bagEdit.Title,
		Category: bagEdit.Category,
		Color:    bagEdit.Color,
		Price:    bagEdit.Price,
		Size:     bagEdit.Size,
	}
	err := s.repo.Create(bag)
	return bag, err
}

// Обновление данных сумки
func (s *BagService) Update(id int, bagEdit *models.BagEdit) (*models.Bag, error) {
	err := s.repo.Update(id, bagEdit)
	if err != nil {
		return nil, err
	}
	return s.GetBagByID(id)
}

// Удаление сумки
func (s *BagService) DeleteBag(bagID int) error {
	return s.repo.Delete(bagID)
}
