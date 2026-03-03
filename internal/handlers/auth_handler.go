package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/finance-tracker/internal/services"
)

type RegisterInput struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// @Summary Kullanıcı kaydı
// @Description Yeni kullanıcı oluşturur
// @Tags auth
// @Accept json
// @Produce json
// @Param input body RegisterInput true "Kullanıcı bilgileri"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(userService *services.UserService, c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.Register(input.UserName, input.Email, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı oluşturulamadı"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kayıt başarılı"})
}

// @Summary Kullanıcı girişi
// @Description Kullanıcı girişi yapar ve token döndürür
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Giriş bilgileri"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(userService *services.UserService, c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := userService.Login(input.Identifier, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
