package database

import (
	logger "client_server/pkg/logger"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	client_config "client_server/pkg/config"
)

type GormDB struct {
	db *gorm.DB
}

func NewGormDB(config client_config.ClientConfig, log logger.Logger) *GormDB {
	return &GormDB{db: generateGormDB(config, log)}
}

func generateGormDB(config client_config.ClientConfig, log logger.Logger) *gorm.DB {
	host := config.Postgres.Host
	user := config.Postgres.User
	password := config.Postgres.Password
	dbName := config.Postgres.DBName
	port := config.Postgres.Port

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		host, user, password, dbName, port)
	log.Debug(dsn)
	var err error
	var db *gorm.DB
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err == nil {
			log.Info("Database connected successfully.")
			return db
		}

		log.Error(fmt.Sprintf("Failed to connect to database (attempt %d/%d)\n", i+1, maxRetries))
		time.Sleep(2 * time.Second)
	}

	panic(fmt.Sprintf("Could not connect to the database after %d attempts: %v", maxRetries, err))
}
