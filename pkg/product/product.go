package product

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ProductModel struct {
}

func (model *ProductModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// h, v := docStyle.GetFrameSize()
		// model.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// if model.list.FilterState() == list.Filtering {
		// 	break
		// }

		switch msg.String() {
		case "ctrl-c", "q":
			return model, tea.Quit

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			// if model, ok := hitems[model.list.Cursor()].(hitem); ok {
			// 	model.Init()
			// 	return model.Model, nil
			// }
		}

	}

	var cmd tea.Cmd
	return model, cmd
}

func (model *ProductModel) View() string {
	return ""
}

func (model *ProductModel) Init() tea.Cmd {
	return nil
}
