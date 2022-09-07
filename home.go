package stream

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

var (
	hitems = []list.Item{
		hitem{
			item{
				"Product Config",
				"product configuration is where you are able to add, update, and remove products",
			},
			new(ProductConfig)},
		hitem{
			item{
				"User Config",
				"user configuration is where you are able to add, update, and remove users",
			},
			new(UserConfig)},
	}
)

type Home struct {
	items    []list.Item
	selected Step
}

func (home Home) Step() string {
	return "Home"
}

func (home *Home) Select(cursor int) error {
	if len(home.items) <= cursor {
		return fmt.Errorf("cursor selected a value that doesn't exist")
	}

	home.selected = home.items[cursor].(hitem).step
	return nil
}

func (home *Home) Init() error {
	home.items = hitems
	return nil
}

type hitem struct {
	item
	step Step
}
