package user

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	pisces "github.com/cryptnode-software/pisces/lib"
	"github.com/cryptnode-software/stream/pkg"
)

type CreateUserModel struct {
	service pisces.AuthService
	inputs  []textinput.Model
	err     error
	focused int
}

const (
	first    = 0
	last     = 1
	password = 2
	username = 3
	email    = 4
)

var (
	inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
)

func (model *CreateUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd = make([]tea.Cmd, len(model.inputs))
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if model.focused == len(model.inputs)-1 {
				ctx := context.Background()

				user := &pisces.User{
					Username: model.inputs[username].Value(),
					Email:    model.inputs[email].Value(),
				}

				model.service.CreateUser(ctx, user, model.inputs[password].Value())

				return model, tea.Quit
			} else {
				model.nextInput()
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return model, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			model.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			model.nextInput()
		}
		for i := range model.inputs {
			model.inputs[i].Blur()
		}
		model.inputs[model.focused].Focus()

	// We handle errors just like any other message
	case error:
		model.err = msg
		return model, nil
	}

	for i := range model.inputs {
		model.inputs[i], cmds[i] = model.inputs[i].Update(msg)
	}
	return model, tea.Batch(cmds...)
}

func (model *CreateUserModel) View() string {
	return fmt.Sprintf(`
		Create User:
		%s
		%s

		%s
		%s

		%s
		%s

		%s
		%s

		%s
		%s
	`,
		inputStyle.Width(20).Render("First Name"),
		model.inputs[first].View(),
		inputStyle.Width(20).Render("Last Name"),
		model.inputs[last].View(),
		inputStyle.Width(20).Render("Password"),
		model.inputs[password].View(),
		inputStyle.Width(20).Render("Username"),
		model.inputs[username].View(),
		inputStyle.Width(20).Render("Email"),
		model.inputs[email].View(),
	)
}

func (model *CreateUserModel) Init() tea.Cmd {
	var err error
	model.service, err = pkg.NewAuthService()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	model.inputs = make([]textinput.Model, 5)

	model.inputs[first] = textinput.New()
	model.inputs[first].Placeholder = "First Name"
	model.inputs[first].Prompt = ""
	model.inputs[first].CharLimit = 20
	model.inputs[first].Width = 20
	model.inputs[first].Focus()

	model.inputs[last] = textinput.New()
	model.inputs[last].Placeholder = "Last Name"
	model.inputs[last].Prompt = ""
	model.inputs[last].CharLimit = 20
	model.inputs[last].Width = 20

	model.inputs[password] = textinput.New()
	model.inputs[password].Placeholder = "Password"
	model.inputs[password].Prompt = ""
	model.inputs[password].CharLimit = 20
	model.inputs[password].Width = 20

	model.inputs[username] = textinput.New()
	model.inputs[username].Placeholder = "Username"
	model.inputs[username].Prompt = ""

	model.inputs[email] = textinput.New()
	model.inputs[email].Placeholder = "Email"
	model.inputs[email].Prompt = ""

	model.focused = first
	model.err = nil

	return textinput.Blink
}

// nextInput focuses the next input field
func (m *CreateUserModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *CreateUserModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
