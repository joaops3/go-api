package config

import (
	"github.com/joaops3/go-api/pkg/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMYSQL() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	
	
	dsn := "root:123@/go_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}