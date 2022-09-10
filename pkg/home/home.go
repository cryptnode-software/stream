package home

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cryptnode-software/stream/pkg/product"
	"github.com/cryptnode-software/stream/pkg/user"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var (
	hitems = []list.Item{
		hitem{
			"Product Config",
			"product configuration is where you are able to add, update, and remove products",
			new(product.ProductModel),
		},
		hitem{
			"User Config",
			"user configuration is where you are able to add, update, and remove users",
			new(user.UserModel),
		},
	}
)

type HomeModel struct {
	list list.Model
}

func (model *HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if model, ok := hitems[model.list.Cursor()].(hitem); ok {
				model.Init()
				return model, nil
			}
		}

	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)
	return model, cmd
}

func (model *HomeModel) View() string {
	return docStyle.Render(model.list.View())
}

func (model *HomeModel) Init() tea.Cmd {
	model.list = list.New(hitems, list.NewDefaultDelegate(), 0, 0)
	return nil
}

type hitem struct {
	title, desc string
	tea.Model
}

func (i hitem) FilterValue() string { return i.title }
func (i hitem) Title() string       { return i.title }
func (i hitem) Description() string { return i.desc }
