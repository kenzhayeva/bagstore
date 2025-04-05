package delivery

import (
	"bagstore/internal/models"
	"bagstore/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Конструктор
func NewBagHandler(service *services.BagService) *BagHandler {
	return &BagHandler{service: service}
}

type BagHandler struct {
	service *services.BagService
}

// Получение всех сумок
func (h *BagHandler) GetAllBags(c *gin.Context) {
	bags, err := h.service.GetAllBags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bags"})
		return
	}
	c.JSON(http.StatusOK, bags)
}

// Получение одной сумки по ID
func (h *BagHandler) GetBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	bag, err := h.service.GetBagByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, bag)
}

// Создание новой сумки
func (h *BagHandler) CreateBag(c *gin.Context) {
	var bagCreate models.BagEdit

	if err := c.ShouldBindJSON(&bagCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newBag, err := h.service.Create(bagCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bag"})
		return
	}

	c.JSON(http.StatusCreated, newBag)
}

// Обновление данных сумки
func (h *BagHandler) UpdateBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	var bagEdit models.BagEdit
	if err := c.ShouldBindJSON(&bagEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedBag, err := h.service.Update(id, &bagEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, updatedBag)
}

// Удаление сумки
func (h *BagHandler) DeleteBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	if err := h.service.DeleteBag(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bag deleted successfully"})
}
