package services

import (
	"github.com/kullaniciadi/finance-tracker/internal/models"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetMonthlySummary(userID, month, year int) (models.MonthlySummary, error) {
	var summary models.MonthlySummary

	s.db.Model(&models.Transactions{}).
		Where("user_id = ? AND EXTRACT(MONTH FROM trans_date) = ? AND EXTRACT(YEAR FROM trans_date) = ? AND ttype = ?", userID, month, year, "gelir").
		Select("COALESCE(SUM(value), 0)").
		Scan(&summary.TotalIncome)

	s.db.Model(&models.Transactions{}).
		Where("user_id = ? AND EXTRACT(MONTH FROM trans_date) = ? AND EXTRACT(YEAR FROM trans_date) = ? AND ttype = ?", userID, month, year, "gider").
		Select("COALESCE(SUM(value), 0)").
		Scan(&summary.TotalExpense)

	summary.NetBalance = summary.TotalIncome - summary.TotalExpense

	return summary, nil
}

func (s *ReportService) GetCategoryExpenses(userID, month, year int) ([]models.CategoryExpense, error) {
	var expenses []models.CategoryExpense

	s.db.Table("transactions").
		Select("categories.category_name, COALESCE(SUM(transactions.value), 0) as total").
		Joins("JOIN transactions_categories ON transactions.trans_id = transactions_categories.trans_id").
		Joins("JOIN categories ON transactions_categories.category_id = categories.category_id").
		Where("transactions.user_id = ? AND EXTRACT(MONTH FROM transactions.trans_date) = ? AND EXTRACT(YEAR FROM transactions.trans_date) = ? AND transactions.ttype = ?", userID, month, year, "gider").
		Group("categories.category_name").
		Scan(&expenses)

	return expenses, nil
}

func (s *ReportService) GetMonthlyComparison(userID, month, year int) (models.MonthlyComparison, error) {
	var comparison models.MonthlyComparison

	current, _ := s.GetMonthlySummary(userID, month, year)
	comparison.CurrentMonth = current

	prevMonth := month - 1
	prevYear := year
	if prevMonth == 0 {
		prevMonth = 12
		prevYear = year - 1
	}

	previous, _ := s.GetMonthlySummary(userID, prevMonth, prevYear)
	comparison.PreviousMonth = previous

	comparison.ExpenseDiff = current.TotalExpense - previous.TotalExpense
	comparison.IncomeDiff = current.TotalIncome - previous.TotalIncome

	return comparison, nil
}

func (s *ReportService) SetBudget(userID, categoryID, month, year int, limit float64) error {
	var budget models.Budget
	result := s.db.Where("user_id = ? AND category_id = ? AND month = ? AND year = ?", userID, categoryID, month, year).First(&budget)

	if result.Error != nil {
		// Yoksa yeni oluştur
		budget = models.Budget{
			UserID:     userID,
			CategoryID: categoryID,
			Limit:      limit,
			Month:      month,
			Year:       year,
		}
		s.db.Create(&budget)
	} else {
		// Varsa güncelle
		budget.Limit = limit
		s.db.Save(&budget)
	}

	return nil
}

func (s *ReportService) GetBudgetStatus(userID, month, year int) ([]models.BudgetStatus, error) {
	var budgets []models.Budget
	s.db.Where("user_id = ? AND month = ? AND year = ?", userID, month, year).Find(&budgets)

	var result []models.BudgetStatus
	for _, budget := range budgets {
		var spent float64
		s.db.Model(&models.Transactions{}).
			Where("user_id = ? AND EXTRACT(MONTH FROM trans_date) = ? AND EXTRACT(YEAR FROM trans_date) = ?", userID, month, year).
			Select("COALESCE(SUM(value), 0)").
			Scan(&spent)

		result = append(result, models.BudgetStatus{
			CategoryID: budget.CategoryID,
			Limit:      budget.Limit,
			Spent:      spent,
			Remaining:  budget.Limit - spent,
			Percentage: (spent / budget.Limit) * 100,
		})
	}

	return result, nil
}
