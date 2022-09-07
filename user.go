package stream

import "github.com/charmbracelet/bubbles/list"

var (
	uitems = []list.Item{
		uitem{
			item{
				"Create User", "create a new user and configure them accordingly",
			},
		},
		uitem{
			item{
				"Create User", "create a new user and configure them accordingly",
			},
		},
	}
)

type uitem struct {
	item
}

type UserConfig struct {
	items []list.Item
}

func (config *UserConfig) Init() {
	config.items = uitems
}

func (config UserConfig) Step() string {
	return "UserConfig"
}

func (config UserConfig) Select(cursor int) error {
	return nil
}
