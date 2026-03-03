package services

import (
	"errors"

	"github.com/kullaniciadi/finance-tracker/internal/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) CreateCategory(userID int, name, categoryType string, parentID *int) error {
	category := models.Category{
		CategoryName: name,
		CategoryType: categoryType,
		ParentID:     parentID,
		UserID:       userID,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return errors.New("kategori oluşturulamadı")
	}

	return nil
}

func (s *CategoryService) GetCategories(userID int) ([]models.Category, error) {
	var categories []models.Category
	if err := s.db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, errors.New("kategoriler getirilemedi")
	}
	return categories, nil
}

func (s *CategoryService) UpdateCategory(userID, categoryID int, name, categoryType string) error {
	var category models.Category
	if err := s.db.Where("category_id = ? AND user_id = ?", categoryID, userID).First(&category).Error; err != nil {
		return errors.New("kategori bulunamadı")
	}

	category.CategoryName = name
	category.CategoryType = categoryType

	if err := s.db.Save(&category).Error; err != nil {
		return errors.New("kategori güncellenemedi")
	}

	return nil
}
func (s *CategoryService) DeleteCategory(userID, categoryID int) error {
	var category models.Category
	if err := s.db.Where("category_id = ? AND user_id = ?", categoryID, userID).First(&category).Error; err != nil {
		return errors.New("kategori bulunamadı")
	}

	if category.IsDefault {
		return errors.New("varsayılan kategori silinemez")
	}

	if err := s.db.Delete(&category).Error; err != nil {
		return errors.New("kategori silinemedi")
	}

	return nil
}
