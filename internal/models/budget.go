package models

type Budget struct {
	BudgetID   int `gorm:"primaryKey"`
	UserID     int
	CategoryID int
	Limit      float64 `gorm:"not null"`
	Month      int
	Year       int
}
