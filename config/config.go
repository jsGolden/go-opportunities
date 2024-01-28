package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}
