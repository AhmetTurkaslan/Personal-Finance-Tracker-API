package models

import "time"

type User struct {
	UserID   int `gorm:"primaryKey"`
	Name     string
	LastName string
	UserName string    `gorm:unique;noy null`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Birthday time.Time `gorm:"not null"`
}

type Category struct {
	CategoryID   int `gorm:"primaryKey"`
	UserID       int
	CategoryName string
	CategoryType string
	ParentID     *int
	IsDefault    bool `gorm:"default:false"`
}
type Transactions_category struct {
	CategoryID int `gorm:"primaryKey"`
	TransID    int `gorm:"primaryKey"`
}

type Transactions struct {
	TransID   int       `gorm:"primaryKey"`
	TransDate time.Time `gorm:"autoCreateTime"`
	UserID    int
	Ttype     string
	Value     float64 `gorm:"not null"`
	Comment   string
}
