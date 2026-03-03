package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kullaniciadi/finance-tracker/config"
	"github.com/kullaniciadi/finance-tracker/internal/handlers"
	"github.com/kullaniciadi/finance-tracker/internal/middleware"
	"github.com/kullaniciadi/finance-tracker/internal/models"
	"github.com/kullaniciadi/finance-tracker/internal/services"

	_ "github.com/kullaniciadi/finance-tracker/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Finance Tracker API
// @version 1.0
// @description Kişisel finans takip uygulaması API'si
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := config.ConnectDB()

	db.AutoMigrate( //tabloları doğru şekilde yapıyor bizim sql kullanmamıza gerek kalmadan
		&models.User{},
		&models.Category{},
		&models.Transactions{},
		&models.Transactions_category{},
		&models.Budget{},
	)
	r := gin.Default()

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	categoryService := services.NewCategoryService(db)

	protected.POST("/categories", func(c *gin.Context) {
		handlers.CreateCategory(categoryService, c)
	})

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	userService := services.NewUserService(db)

	r.POST("/register", func(c *gin.Context) {
		handlers.Register(userService, c)
	})

	r.POST("/login", func(c *gin.Context) {
		handlers.Login(userService, c)
	})
	protected.GET("/categories", func(c *gin.Context) {
		handlers.GetCategories(categoryService, c)
	})

	protected.PUT("/categories/:id", func(c *gin.Context) {
		handlers.UpdateCategory(categoryService, c)
	})

	protected.DELETE("/categories/:id", func(c *gin.Context) {
		handlers.DeleteCategory(categoryService, c)
	})
	transactionService := services.NewTransactionService(db)

	protected.POST("/transactions", func(c *gin.Context) {
		handlers.CreateTransaction(transactionService, c)
	})

	protected.GET("/transactions", func(c *gin.Context) {
		handlers.GetTransactions(transactionService, c)
	})

	protected.PUT("/transactions/:id", func(c *gin.Context) {
		handlers.UpdateTransaction(transactionService, c)
	})

	protected.DELETE("/transactions/:id", func(c *gin.Context) {
		handlers.DeleteTransaction(transactionService, c)
	})
	reportService := services.NewReportService(db)

	protected.GET("/report/summary", func(c *gin.Context) {
		handlers.GetMonthlySummary(reportService, c)
	})

	protected.GET("/report/categories", func(c *gin.Context) {
		handlers.GetCategoryExpenses(reportService, c)
	})

	protected.GET("/report/comparison", func(c *gin.Context) {
		handlers.GetMonthlyComparison(reportService, c)
	})

	protected.POST("/report/budget", func(c *gin.Context) {
		handlers.SetBudget(reportService, c)
	})

	protected.GET("/report/budget", func(c *gin.Context) {
		handlers.GetBudgetStatus(reportService, c)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + os.Getenv("PORT"))
}
