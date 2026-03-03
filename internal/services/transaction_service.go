package services

import (
	"errors"

	"github.com/kullaniciadi/finance-tracker/internal/models"
	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) CreateTransaction(userID int, ttype string, value float64, comment string) error {
	transaction := models.Transactions{
		UserID:  userID,
		Ttype:   ttype,
		Value:   value,
		Comment: comment,
	}

	if err := s.db.Create(&transaction).Error; err != nil {
		return errors.New("işlem oluşturulamadı")
	}

	return nil
}

func (s *TransactionService) GetTransactions(userID int) ([]models.Transactions, error) {
	var transactions []models.Transactions
	if err := s.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, errors.New("işlemler getirilemedi")
	}
	return transactions, nil
}

func (s *TransactionService) DeleteTransaction(userID int, transID int) error {
	var transaction models.Transactions
	if err := s.db.Where("trans_id = ? AND user_id = ?", transID, userID).First(&transaction).Error; err != nil {
		return errors.New("işlem bulunamadı")
	}
	if err := s.db.Delete(&transaction).Error; err != nil {
		return errors.New("işlem silinemedi")
	}
	return nil
}

func (s *TransactionService) UpdateTransaction(userID int, trans_id int, value float64, ttype string, comment string) error {
	var transaction models.Transactions
	if err := s.db.Where("trans_id=? AND user_id =?", trans_id, userID).First(&transaction).Error; err != nil {
		return errors.New("İşlem bulunumadı")
	}
	transaction.Value = value
	transaction.Comment = comment
	transaction.Ttype = ttype

	if err := s.db.Save(&transaction).Error; err != nil {
		return errors.New("Değişiklikler kaydedilemedi")
	}
	return nil
}
