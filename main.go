package main

import (
	"fmt"

	"github.com/celio001/api-golang-opportunities.git/config"
	"github.com/celio001/api-golang-opportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Inicialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config inicialization error: %v", err)
		fmt.Println(err)
	}
	// Inicialize Router
	router.Initialize()
}
