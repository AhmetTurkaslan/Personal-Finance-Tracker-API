package models

type MonthlySummary struct {
	TotalIncome  float64
	TotalExpense float64
	NetBalance   float64
}

type CategoryExpense struct {
	CategoryName string
	Total        float64
}

type MonthlyComparison struct {
	CurrentMonth  MonthlySummary
	PreviousMonth MonthlySummary
	ExpenseDiff   float64
	IncomeDiff    float64
}

type BudgetStatus struct {
	CategoryID int
	Limit      float64
	Spent      float64
	Remaining  float64
	Percentage float64
}

type BudgetInput struct {
	CategoryID int     `json:"category_id" binding:"required"`
	Limit      float64 `json:"limit" binding:"required"`
	Month      int     `json:"month" binding:"required"`
	Year       int     `json:"year" binding:"required"`
}
