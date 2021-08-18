package main

import (
	"github.com/DavinciUI/backend/config"
	"github.com/DavinciUI/backend/router"
)

func main() {
	configuration, configErr := config.Load(".env")
	if configErr != nil {
		panic(configErr)
	}

	_, routerErr := router.CreateRouter(*configuration)
	if routerErr != nil {
		panic(routerErr)
	}
}