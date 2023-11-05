package model

import "gorm.io/gorm"

type Config struct {
	DBUsername string `yaml:"DBUsername"`
	DBPassword string `yaml:"DBPassword"`
	DBName     string `yaml:"DBName"`
	DBHost     string `yaml:"DBHost"`
	DBPort     string `yaml:"DBPort"`
}

type DbArticle struct {
	gorm.Model
	Title   string `gorm:"type:text;notnull"`
	Content string `gorm:"type:text;notnull"`
}
