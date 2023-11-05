package db

import (
	m "WebBlog/internal/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect(configPath string) (*gorm.DB, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&m.DbArticle{}); err != nil {
		return nil, err
	}
	return db, nil
}

func loadConfig(configPath string) (*m.Config, error) {
	var config m.Config

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
