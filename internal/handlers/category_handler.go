package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/finance-tracker/internal/services"
)

type CategoryInput struct {
	CategoryName string `json:"category_name" binding:"required"`
	Type         string `json:"type" binding:"required"`
	ParentID     *int   `json:"parent_id"`
}

// @Summary Kategori oluştur
// @Description Yeni kategori oluşturur
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body CategoryInput true "Kategori bilgileri"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func CreateCategory(categoryService *services.CategoryService, c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := int(c.MustGet("user_id").(float64))

	if err := categoryService.CreateCategory(userID, input.CategoryName, input.Type, input.ParentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kategori oluşturuldu"})
}

// @Summary Kategorileri listele
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /categories [get]
func GetCategories(categoryService *services.CategoryService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))

	categories, err := categoryService.GetCategories(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// @Summary Kategori güncelle
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Kategori ID"
// @Param input body CategoryInput true "Kategori bilgileri"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [put]
func UpdateCategory(categoryService *services.CategoryService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	categoryID := c.Param("id")

	var input struct {
		CategoryName string `json:"category_name" binding:"required"`
		Type         string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := 0
	fmt.Sscanf(categoryID, "%d", &id)

	if err := categoryService.UpdateCategory(userID, id, input.CategoryName, input.Type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori güncellendi"})
}

// @Summary Kategori sil
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Kategori ID"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(categoryService *services.CategoryService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	categoryID := c.Param("id")

	id := 0
	fmt.Sscanf(categoryID, "%d", &id)

	if err := categoryService.DeleteCategory(userID, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori silindi"})
}
