package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cryptnode-software/stream/pkg/home"
)

func main() {
	model := new(Model).Start()
	if err := tea.NewProgram(model).Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Step interface {
	Select(cursor int) error
	Step() string
}

type Model struct {
	history []tea.Model
	tea.Model
}

func (model *Model) Init() tea.Cmd {
	model.Model.Init()
	return nil
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model.Model.Update(msg)
}

func (model *Model) View() string {
	return model.Model.View()
}

func (model *Model) Start() *Model {
	model.Model = new(home.HomeModel)
	return model
}
