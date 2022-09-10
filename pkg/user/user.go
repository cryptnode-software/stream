package user

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	uitems = []list.Item{

		item{
			"Create User", "create a new user and configure them accordingly",
		},
	}
)

var userdocStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) FilterValue() string { return i.title }
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }

type UserModel struct {
	list list.Model
}

func (model *UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := userdocStyle.GetFrameSize()
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
			return model, nil
			// return uitems[model.list.Cursor()].(uitem).Model, nil
		}

	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)
	return model, cmd

}

func (model *UserModel) View() string {
	return userdocStyle.Render(model.list.View())
}

func (model *UserModel) Init() tea.Cmd {
	model.list = list.New(uitems, list.NewDefaultDelegate(), 20, 20)
	return nil
}
