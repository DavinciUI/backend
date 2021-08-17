package entity

import "github.com/DavinciUI/backend/code"

func Create(
	code code.Code,
	title []string,
	options MenuOptions,
	rows []MenuItem,
	actions map[string][]string) *Menu {
	return &Menu{code, title, options, rows, actions}
}

func CreateOptions(updateTicks int, papiSupport bool, hidePlayerInventory bool) *MenuOptions {
	return &MenuOptions{updateTicks, papiSupport, hidePlayerInventory}
}

func CreateItem(material string, name string, lore []string, meta ItemMeta, actions map[ClickType]string) *MenuItem {
	return &MenuItem{material, name, lore, meta, actions}
}

type Menu struct {

	CodeId code.Code `json:"code"`
	Title []string `json:"title"`
	Options MenuOptions `json:"options"`
	Rows []MenuItem `json:"rows"`
	Actions map[string][]string `json:"actions"`

}

// GetCode Implementing Entity interface in this struct you can get the Code
func (menu Menu) GetCode() code.Code {
	return menu.CodeId
}

type MenuOptions struct {

	TitleUpdate int `json:"title-update"`
	PlaceholderSupport bool `json:"placeholder-api-support"`
	HidePlayerInventory bool `json:"hide-player-inventory"`

}

type MenuItem struct {

	Material string `json:"type"`
	Name string `json:"display-name"`
	Lore []string `json:"lore"`
	Meta ItemMeta `json:"meta"`
	ActionReference map[ClickType]string `json:"actions"`

}

type ClickType int

const (
	LEFT_CLICK ClickType = 1 << iota
	RIGHT_CLICK
)

type ItemMeta struct {
	ModelData int
}
