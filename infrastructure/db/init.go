package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DEFAULT_HOST = "localhost"
)

func InitDB() *gorm.DB {
	dsn := generateConnectionString()
	logLevel := logger.Info
	if os.Getenv("POSTGRES_LOG_LEVEL") == "silent" {
		logLevel = logger.Silent
	} else if os.Getenv("POSTGRES_LOG_LEVEL") == "error" {
		logLevel = logger.Error
	} else if os.Getenv("POSTGRES_LOG_LEVEL") == "warn" {
		logLevel = logger.Warn
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	DbExceptionHandle(err)
	//Migrate(db)
	//migration function
	return db
}

func generateConnectionString() string {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARNING] %v", ".env file not found. Make sure to load environment variables from somewhere else. Look into .env.example for reference on environemnt variables")
		log.Printf("[WARNING] %v", err)
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("SSL_MODE")
	if host == "" {
		host = DEFAULT_HOST
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, db, port, sslmode)
}

func DbExceptionHandle(err error) {
	if err != nil {
		panic(err)
	}
}
