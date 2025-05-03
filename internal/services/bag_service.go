package services

import (
	"bagstore/internal/models"
	"github.com/pkg/errors"
)

var (
	ErrBagNotFound = errors.New("bag not found")
)

// Интерфейс репозитория сумок
type BagRepository interface {
	GetAll(userID uint) ([]models.Bag, error)
	GetByID(userID uint, bagID uint) (*models.Bag, error)
	Create(bag *models.Bag) error
	Update(userID uint, bagID uint, bag *models.BagEdit) error
	Delete(userID uint, bagID uint) error
}

// Структура сервиса сумок
type BagService struct {
	repo BagRepository
}

func NewBagService(repo BagRepository) *BagService {
	return &BagService{repo: repo}
}

func (s *BagService) GetAllBags(userID uint) ([]models.Bag, error) {
	return s.repo.GetAll(userID)
}

func (s *BagService) GetBagByID(userID uint, bagID uint) (*models.Bag, error) {
	return s.repo.GetByID(userID, bagID)
}

func (s *BagService) CreateBag(bag *models.Bag) error {
	return s.repo.Create(bag)
}

func (s *BagService) UpdateBag(userID uint, bagID uint, bagEdit *models.BagEdit) error {
	return s.repo.Update(userID, bagID, bagEdit)
}

func (s *BagService) DeleteBag(userID uint, bagID uint) error {
	return s.repo.Delete(userID, bagID)
}
