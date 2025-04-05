//package authorization
//
//import (
//	"bagstore/internal/db"
//	"bagstore/internal/models"
//	"github.com/gin-gonic/gin"
//	"golang.org/x/crypto/bcrypt"
//	"net/http"
//)
//
//func Login(c *gin.Context) {
//	var req struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
//		return
//	}
//
//	var u models.User
//	result := db.DB.Where("username = ?", req.Username).First(&u)
//	if result.Error != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
//		return
//	}
//
//	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
//		return
//	}
//
//	token, err := GenerateJWT(u.ID)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"token": token})
//}
//
//func Register(c *gin.Context) {
//	var req struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя пользователя и пароль обязательны"})
//		return
//	}
//
//	var existing models.User
//	if err := db.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
//		c.JSON(http.StatusConflict, gin.H{"error": "Имя пользователя уже занято"})
//		return
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке пароля"})
//		return
//	}
//
//	newUser := models.User{
//		Username: req.Username,
//		Password: string(hashedPassword),
//	}
//
//	if err := db.DB.Create(&newUser).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
//}
//func Me(c *gin.Context) {
//	userID, exists := c.Get("userID")
//	if !exists {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in context"})
//		return
//	}
//
//	var u models.User
//	if err := db.DB.First(&u, userID).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"id":       u.ID,
//		"username": u.Username,
//	})
//}

package authorization

import (
	"bagstore/internal/db"
	"bagstore/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	var u models.User
	result := db.DB.Where("username = ?", req.Username).First(&u)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	token, err := GenerateJWT(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя пользователя и пароль обязательны"})
		return
	}

	var existing models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Имя пользователя уже занято"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке пароля"})
		return
	}
	// admin
	var count int64
	DB.Model(&models.User{}).Count(&count)

	role := "user"
	if count == 0 {
		role = "admin" // Первый — админ
	}

	newUser := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}
func Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in context"})
		return
	}

	var u models.User
	if err := db.DB.First(&u, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       u.ID,
		"username": u.Username,
	})
}
