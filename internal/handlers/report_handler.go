package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/finance-tracker/internal/models"
	"github.com/kullaniciadi/finance-tracker/internal/services"
)

// @Summary Aylık özet
// @Tags report
// @Produce json
// @Security BearerAuth
// @Param month query int true "Ay"
// @Param year query int true "Yıl"
// @Success 200 {object} map[string]interface{}
// @Router /report/summary [get]
func GetMonthlySummary(reportService *services.ReportService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	summary, err := reportService.GetMonthlySummary(userID, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

// @Summary Kategori bazında harcamalar
// @Tags report
// @Produce json
// @Security BearerAuth
// @Param month query int true "Ay"
// @Param year query int true "Yıl"
// @Success 200 {object} map[string]interface{}
// @Router /report/categories [get]
func GetCategoryExpenses(reportService *services.ReportService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	expenses, err := reportService.GetCategoryExpenses(userID, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

// @Summary Aylık karşılaştırma
// @Tags report
// @Produce json
// @Security BearerAuth
// @Param month query int true "Ay"
// @Param year query int true "Yıl"
// @Success 200 {object} map[string]interface{}
// @Router /report/comparison [get]
func GetMonthlyComparison(reportService *services.ReportService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	comparison, err := reportService.GetMonthlyComparison(userID, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comparison": comparison})
}

// @Summary Bütçe limiti belirle
// @Tags report
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.BudgetInput true "Bütçe bilgileri"
// @Success 201 {object} map[string]string
// @Router /report/budget [post]
func SetBudget(reportService *services.ReportService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	var input models.BudgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := reportService.SetBudget(userID, input.CategoryID, input.Month, input.Year, input.Limit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Bütçe oluşturuldu"})
}

// @Summary Bütçe durumu
// @Tags report
// @Produce json
// @Security BearerAuth
// @Param month query int true "Ay"
// @Param year query int true "Yıl"
// @Success 200 {object} map[string]interface{}
// @Router /report/budget [get]
func GetBudgetStatus(reportService *services.ReportService, c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	status, err := reportService.GetBudgetStatus(userID, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budget_status": status})
}
