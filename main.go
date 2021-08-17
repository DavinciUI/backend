package main

import (
	"fmt"
	"github.com/DavinciUI/backend/code"
	"github.com/DavinciUI/backend/entity"
	"github.com/DavinciUI/backend/repository"
)

func main() {
	cache := repository.NewCachedRepository()

	options := entity.CreateOptions(20, true, false)

	actionReference := make(map[entity.ClickType]string)
	actionReference[entity.LeftClick] = "test"

	itemExample := entity.CreateItem("STONE", "test", make([]string, 2), entity.ItemMeta{ModelData: 2}, actionReference)

	items := make([]entity.MenuItem, 1)
	items[0] = *itemExample

	title := []string{""}
	title[0] = "Test"

	rows := []entity.MenuItem{*itemExample}

	actions := make(map[string][]string)
	actions["test"] = []string{"testAction"}

	menu := entity.Create(*code.GenerateNewCode(), title, *options, rows, actions)

	err := cache.Save(menu)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(cache.FindAll()[0])
}