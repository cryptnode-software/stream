package user

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	pisces "github.com/cryptnode-software/pisces/lib"
	"github.com/cryptnode-software/stream/pkg"
)

type CreateUserModel struct {
	service pisces.AuthService
}

func (model *CreateUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	return model, cmd
}

func (model *CreateUserModel) View() string {
	return ""
}

func (model *CreateUserModel) Init() tea.Cmd {
	var err error
	model.service, err = pkg.NewAuthService()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return nil
}
