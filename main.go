package main

import (
	"fmt"
	"os"
	"strconv"

	"gokana/internal/game"
	"gokana/internal/model"
	"gokana/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

// teaModel wraps our model to implement tea.Model interface
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
	startLevel := 1

	// Check for command-line argument for starting level
	if len(os.Args) > 1 {
		var err error
		startLevel, err = strconv.Atoi(os.Args[1])
		if err != nil || startLevel < 1 {
			fmt.Println("Usage: gokana [level]")
			fmt.Println("  level: Starting level (default: 1)")
			os.Exit(1)
		}
	}

	initialModel := game.InitialModel(startLevel)
	p := tea.NewProgram(teaModel{m: initialModel})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
