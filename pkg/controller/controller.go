package controller

import (
	"github.com/joaops3/go-api/pkg/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	Db *gorm.DB
)

func InitController(){
	logger = config.GetLogger("controller")
	Db = config.GetDb()
}