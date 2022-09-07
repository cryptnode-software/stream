package stream

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	pisces "github.com/cryptnode-software/pisces/lib"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type services struct {
	User *pisces.AuthService
}

type Step interface {
	Select(cursor int) error
	Step() string
}

type Model struct {
	list   list.Model
	Step   Step
	cursor int
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		model.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if model.list.FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl-c", "q":
			return model, tea.Quit

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			model.Step.Select(model.list.Cursor())
			model.Next()
		}

	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)
	return model, cmd

}

type item struct {
	title, desc string
}

func (i item) FilterValue() string { return i.title }
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }

func (model Model) View() string {
	return docStyle.Render(model.list.View())
}

func (model Model) Start() Model {
	home := new(Home)
	home.Init()

	model.list = list.New(home.items, list.NewDefaultDelegate(), 0, 0)
	model.Step = home

	return model
}

func (model *Model) Next() {
	switch step := model.Step.(type) {
	case *Home:
		switch selected := step.selected.(type) {
		case *UserConfig:
			selected.Init()
			model.list = list.New(selected.items, list.NewDefaultDelegate(), 0, 0)
			model.Step = step.selected
		}
	}
}
