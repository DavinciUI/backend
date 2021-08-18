package main

import (
	"fmt"
	"github.com/DavinciUI/backend/config"
	"github.com/DavinciUI/backend/router"
	"net/http"
	"time"
)

func main() {
	configuration, configErr := config.Load(".env")
	if configErr != nil {
		panic(configErr)
	}

	go test_rateList()
	go test_rateList()
	go test_rateList()
	go test_rateList()
	go test_rateList()

	_, routerErr := router.CreateRouter(*configuration)
	if routerErr != nil {
		panic(routerErr)
	}
}

func test_rateList() {
	fmt.Println("GOROUTINE")
	time.Sleep(time.Second * 3)

	response, err := http.Get("http://127.0.0.1:7777/home")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.StatusCode)
}