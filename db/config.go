package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getDbConfig() *DbConfig {
	return &DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "root",
		Database: "test",
	}
}

func GetDbClient() *gorm.DB {
	config := getDbConfig()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dsn := fmt.Sprintf(`
      host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Seoul",
      config.Host, config.User, config.Password, config.Database, config.Port
  `)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Database not connected error")
	}
	return db
}
