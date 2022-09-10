package user

import tea "github.com/charmbracelet/bubbletea"

type CreateUserModel struct {
}

func (model *CreateUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	// model.list, cmd = model.list.Update(msg)
	return model, cmd
}

func (model *CreateUserModel) View() string {
	return ""
}

func (model *CreateUserModel) Init() tea.Cmd {
	return nil
}
