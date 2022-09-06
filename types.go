package stream

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	selected map[int]struct{}
	choices  []string
	cursor   int
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return model, tea.Quit
			// The "up" and "k" keys move the cursor up
		case "up", "k":
			if model.cursor > 0 {
				model.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if model.cursor < len(model.choices)-1 {
				model.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := model.selected[model.cursor]
			if ok {
				delete(model.selected, model.cursor)
			} else {
				model.selected[model.cursor] = struct{}{}
			}
		}
	}

	return model, nil
}

func (model Model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range model.choices {
		cursor := " "

		if model.cursor == 1 {
			cursor = ">"
		}

		checked := " "

		if _, ok := model.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func (model Model) Start() Model {
	model.choices = []string{
		"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
	}
	model.selected = make(map[int]struct{})
	return model
}
