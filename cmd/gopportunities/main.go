package main

import (
	"github.com/ryan-albuquerque/gopportunities-api/api/http/router"
	"github.com/ryan-albuquerque/gopportunities-api/config"
	dbConfig "github.com/ryan-albuquerque/gopportunities-api/internal/config"
)

const (
	context = "MAIN"
)

func main() {
	err := dbConfig.Init()
	if err != nil {
		config.LogError(context, "Error during establish DB connection: %v", err)
		return
	}
	router.Initialize()
}