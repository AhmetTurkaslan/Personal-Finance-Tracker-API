package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kullaniciadi/finance-tracker/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}
func (s *UserService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		UserName: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return err
	}

	// Varsayılan kategorileri ekle
	defaultCategories := []struct {
		Name string
		Type string
	}{
		{"Market", "gider"},
		{"Ulaşım", "gider"},
		{"Hobi", "gider"},
		{"Alışveriş", "gider"},
		{"Sağlık", "gider"},
		{"Bağış", "gider"},
		{"Maaş", "gelir"},
		{"Ek Gelir", "gelir"},
	}

	for _, cat := range defaultCategories {
		s.db.Create(&models.Category{
			CategoryName: cat.Name,
			CategoryType: cat.Type,
			UserID:       user.UserID,
			IsDefault:    true,
		})
	}

	return nil
}

func (s *UserService) Login(identifier, password string) (string, error) {
	var user models.User
	if strings.Contains(identifier, "@") {
		s.db.Where("email = ?", identifier).First(&user)
	} else {
		s.db.Where("user_name = ?", identifier).First(&user)
	}

	if user.UserID == 0 {
		return "", errors.New("kullanıcı bulunamadı")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("şifre yanlış")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("token oluşturulamadı")
	}

	return tokenString, nil
}
