package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminHandler struct {
	DB *gorm.DB
}

// Удалить всех пользователей
func (h *AdminHandler) DeleteAllUsers(c *gin.Context) {
	if err := h.DB.Exec("DELETE FROM users").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении всех пользователей"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Все пользователи удалены"})
}

// Удалить одного пользователя по ID
func (h *AdminHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Exec("DELETE FROM users WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удалён"})
}
