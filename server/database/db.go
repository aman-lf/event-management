package database

import (
	"github.com/aman-lf/event-management/config"
	"github.com/aman-lf/event-management/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func ConnectDB() (*gorm.DB, error) {
	dbUser := config.Cfg.DBUser
	dbPassword := config.Cfg.DBPassword
	dbName := config.Cfg.DBName
	dbHost := config.Cfg.DBHost
	dbPort := config.Cfg.DBPort

	dsn := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"

	// Initialize GORM DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.Participant{})
	db.AutoMigrate(&model.Expense{})
	db.AutoMigrate(&model.Activity{})

	DB = db

	return db, nil
}
