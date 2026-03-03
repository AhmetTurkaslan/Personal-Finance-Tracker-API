package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/finance-tracker/internal/services"
)

type TransactionInput struct {
	Ttype   string  `json:"ttype" binding:"required"`
	Value   float64 `json:"value" binding:"required"`
	Comment string  `json:"comment"`
}

// @Summary İşlem oluştur
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body TransactionInput true "İşlem bilgileri"
// @Success 201 {object} map[string]string
// @Router /transactions [post]
func CreateTransaction(transactionService *services.TransactionService, c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := int(c.MustGet("user_id").(float64))
	if err := transactionService.CreateTransaction(userID, input.Ttype, input.Value, input.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "İşlem oluşturuldu"})
}

// @Summary İşlemleri listele
// @Tags transactions
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /transactions [get]
func GetTransactions(transactionService *services.TransactionService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))

	transactions, err := transactionService.GetTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// @Summary İşlem güncelle
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "İşlem ID"
// @Param input body TransactionInput true "İşlem bilgileri"
// @Success 200 {object} map[string]string
// @Router /transactions/{id} [put]
func UpdateTransaction(transactionService *services.TransactionService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	transID := c.Param("id")

	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := 0
	fmt.Sscanf(transID, "%d", &id)

	if err := transactionService.UpdateTransaction(userID, id, input.Value, input.Ttype, input.Comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "işlem güncellendi"})
}

// @Summary İşlem sil
// @Tags transactions
// @Produce json
// @Security BearerAuth
// @Param id path int true "İşlem ID"
// @Success 200 {object} map[string]string
// @Router /transactions/{id} [delete]
func DeleteTransaction(transactionService *services.TransactionService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	transID := c.Param("id")

	id := 0
	fmt.Sscanf(transID, "%d", &id)

	if err := transactionService.DeleteTransaction(userID, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "işlem silindi"})
}
