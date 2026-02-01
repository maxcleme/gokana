package main

import (
	"fmt"
	"os"

	"gokana/internal/game"
	"gokana/internal/model"
	"gokana/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type teaModel struct {
	m *model.Model
}

func (t teaModel) Init() tea.Cmd {
	return game.Init()
}

func (t teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := game.Update(t.m, msg)
	t.m = updatedModel
	return t, cmd
}

func (t teaModel) View() string {
	return ui.View(t.m)
}

func main() {
	initialModel := game.InitialModel()
	p := tea.NewProgram(teaModel{m: initialModel})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
