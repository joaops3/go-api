package main

import (
	"github.com/joaops3/go-api/pkg/config"
	"github.com/joaops3/go-api/pkg/router"
)

var (
	logger config.Logger
)
func main() {
	logger = *config.GetLogger("main")
	config.LoadEnvFile()
	err := config.Init()
	if err != nil {
		logger.Errorf("Config inicialization err %v", err)
		panic(err.Error())
		
	}
	router.Initialize()
}