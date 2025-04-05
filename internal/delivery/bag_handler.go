<<<<<<< HEAD
=======
// package delivery
//
// import (
//
//	"bagstore/internal/models"
//	"bagstore/internal/services"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//
// )
//
//	func NewBagHandler(service *services.BagService) *BagHandler {
//		return &BagHandler{service: service}
//	}
//
//	type BagHandler struct {
//		service *services.BagService
//	}
//
//	func (h *BagHandler) GetAllBags(c *gin.Context) {
//		bags, err := h.service.GetAllBags()
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bags"})
//			return
//		}
//		c.JSON(http.StatusOK, bags)
//	}
//
//	func (h *BagHandler) GetBag(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
//			return
//		}
//
//		bag, err := h.service.GetBagByID(id)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
//			return
//		}
//
//		c.JSON(http.StatusOK, bag)
//	}
//
//	func (h *BagHandler) CreateBag(c *gin.Context) {
//		var bagCreate models.BagEdit
//
//		if err := c.ShouldBindJSON(&bagCreate); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//			return
//		}
//
//		newBag, err := h.service.Create(bagCreate)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bag"})
//			return
//		}
//
//		c.JSON(http.StatusCreated, newBag)
//	}
//
//	func (h *BagHandler) UpdateBag(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
//			return
//		}
//
//		var bagEdit models.BagEdit
//		if err := c.ShouldBindJSON(&bagEdit); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//			return
//		}
//
//		updatedBag, err := h.service.Update(id, &bagEdit)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
//			return
//		}
//
//		c.JSON(http.StatusOK, updatedBag)
//	}
//
//	func (h *BagHandler) DeleteBag(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
//			return
//		}
//
//		if err := h.service.DeleteBag(uint(id)); err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{"message": "Bag deleted successfully"})
//	}
>>>>>>> be3f14a (Initial commit)
package delivery

import (
	"bagstore/internal/models"
	"bagstore/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

<<<<<<< HEAD
// Конструктор
func NewBagHandler(service *services.BagService) *BagHandler {
	return &BagHandler{service: service}
}

=======
>>>>>>> be3f14a (Initial commit)
type BagHandler struct {
	service *services.BagService
}

<<<<<<< HEAD
// Получение всех сумок
func (h *BagHandler) GetAllBags(c *gin.Context) {
	bags, err := h.service.GetAllBags()
=======
func NewBagHandler(service *services.BagService) *BagHandler {
	return &BagHandler{service: service}
}

func getUserID(c *gin.Context) (uint, bool) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return 0, false
	}
	return uid.(uint), true
}

func (h *BagHandler) GetAllBags(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	bags, err := h.service.GetAllBags(userID)
>>>>>>> be3f14a (Initial commit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bags"})
		return
	}
	c.JSON(http.StatusOK, bags)
}

<<<<<<< HEAD
// Получение одной сумки по ID
func (h *BagHandler) GetBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
=======
func (h *BagHandler) GetBag(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	bagID, err := strconv.Atoi(c.Param("id"))
>>>>>>> be3f14a (Initial commit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

<<<<<<< HEAD
	bag, err := h.service.GetBagByID(id)
=======
	bag, err := h.service.GetBagByID(userID, uint(bagID))
>>>>>>> be3f14a (Initial commit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, bag)
}

<<<<<<< HEAD
// Создание новой сумки
func (h *BagHandler) CreateBag(c *gin.Context) {
	var bagCreate models.BagEdit

	if err := c.ShouldBindJSON(&bagCreate); err != nil {
=======
func (h *BagHandler) CreateBag(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var bagEdit models.BagEdit
	if err := c.ShouldBindJSON(&bagEdit); err != nil {
>>>>>>> be3f14a (Initial commit)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

<<<<<<< HEAD
	newBag, err := h.service.Create(bagCreate)
	if err != nil {
=======
	bag := &models.Bag{
		Title:    bagEdit.Title,
		Category: bagEdit.Category,
		Color:    bagEdit.Color,
		Price:    bagEdit.Price,
		Size:     bagEdit.Size,
		UserID:   userID,
	}

	if err := h.service.Create(bag); err != nil {
>>>>>>> be3f14a (Initial commit)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bag"})
		return
	}

<<<<<<< HEAD
	c.JSON(http.StatusCreated, newBag)
}

// Обновление данных сумки
func (h *BagHandler) UpdateBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
=======
	c.JSON(http.StatusCreated, bag)
}

func (h *BagHandler) UpdateBag(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	bagID, err := strconv.Atoi(c.Param("id"))
>>>>>>> be3f14a (Initial commit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

	var bagEdit models.BagEdit
	if err := c.ShouldBindJSON(&bagEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

<<<<<<< HEAD
	updatedBag, err := h.service.Update(id, &bagEdit)
	if err != nil {
=======
	if err := h.service.Update(userID, uint(bagID), &bagEdit); err != nil {
>>>>>>> be3f14a (Initial commit)
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, updatedBag)
}

// Удаление сумки
func (h *BagHandler) DeleteBag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
=======
	c.JSON(http.StatusOK, gin.H{"message": "Bag updated successfully"})
}

func (h *BagHandler) DeleteBag(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	bagID, err := strconv.Atoi(c.Param("id"))
>>>>>>> be3f14a (Initial commit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bag ID"})
		return
	}

<<<<<<< HEAD
	if err := h.service.DeleteBag(id); err != nil {
=======
	if err := h.service.Delete(userID, uint(bagID)); err != nil {
>>>>>>> be3f14a (Initial commit)
		c.JSON(http.StatusNotFound, gin.H{"error": "Bag not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bag deleted successfully"})
}
