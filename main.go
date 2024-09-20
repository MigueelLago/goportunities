package main

import (
	"github.com/MigueelLago/goportunities/config"
	"github.com/MigueelLago/goportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Inicializa as configs
	err := config.Init()

	if err != nil {
		logger.ErrorF("Erro na inicialização: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
