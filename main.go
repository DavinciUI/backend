package main

import (
	"fmt"
	"github.com/DavinciUI/backend/repository"
)

func main() {
	cache := repository.NewCachedRepository()

	fmt.Println(cache.FindAll())
}