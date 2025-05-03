package delivery

import (
	"bagstore/internal/models"
	"bagstore/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BagHandler struct {
	service *services.BagService
}

func NewBagHandler(service *services.BagService) *BagHandler {
	return &BagHandler{service: service}
}

// Получить все сумки
func (h *BagHandler) GetAllBags(c *gin.Context) {
	userID := c.GetUint("userID") // предполагаем, что мидлварь добавила userID
	bags, err := h.service.GetAllBags(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bags"})
		return
	}
	c.JSON(http.StatusOK, bags)
}

// Получить одну сумку
func (h *BagHandler) GetBag(c *gin.Context) {
	userID := c.GetUint("userID")
	bagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	bag, err := h.service.GetBagByID(userID, uint(bagID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}
	c.JSON(http.StatusOK, bag)
}

// Создать сумку
func (h *BagHandler) CreateBag(c *gin.Context) {
	userID := c.GetUint("userID")

	var bagEdit models.BagEdit
	if err := c.ShouldBindJSON(&bagEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newBag := &models.Bag{
		Title:    bagEdit.Title,
		Category: bagEdit.Category,
		Color:    bagEdit.Color,
		Price:    bagEdit.Price,
		Size:     bagEdit.Size,
		UserID:   userID,
	}

	if err := h.service.CreateBag(newBag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bag"})
		return
	}

	c.JSON(http.StatusCreated, newBag)
}

// Обновить сумку
func (h *BagHandler) UpdateBag(c *gin.Context) {
	userID := c.GetUint("userID")
	bagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	var bagEdit models.BagEdit
	if err := c.ShouldBindJSON(&bagEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.service.UpdateBag(userID, uint(bagID), &bagEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bag updated successfully"})
}

// Удалить сумку
func (h *BagHandler) DeleteBag(c *gin.Context) {
	userID := c.GetUint("userID")
	bagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	err = h.service.DeleteBag(uint(userID), uint(bagID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bag deleted successfully"})
}
