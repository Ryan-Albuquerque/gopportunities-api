package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ryan-albuquerque/gopportunities-api/config"
	"github.com/ryan-albuquerque/gopportunities-api/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("no .env file found. Using system environment variables")
	}

	// Read environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")
	parseTime := os.Getenv("DB_PARSE_TIME")
	loc := os.Getenv("DB_LOC")

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password, host, port, dbName, charset, parseTime, loc)

	// Connect to MySQL
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Opening{}); err != nil {
		return fmt.Errorf("failed to migrate tables to database: %w", err)
	}

	config.LogInfo("DB", "Database connection established successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
