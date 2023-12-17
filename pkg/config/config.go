package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeMYSQL()

	if err != nil {
		return fmt.Errorf("error initializing mysql: %v", err)
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}