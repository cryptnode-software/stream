package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cryptnode-software/stream"
)

func main() {
	model := new(stream.Model).Start()
	if err := tea.NewProgram(model).Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
