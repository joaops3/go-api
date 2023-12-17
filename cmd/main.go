package main

import (
	"github.com/joaops3/go-api/config"
	"github.com/joaops3/go-api/router"
)

var (
	logger config.Logger
)
func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config inicialization err %v", err)
		panic(err.Error())
		
	}
	router.Initialize()
}