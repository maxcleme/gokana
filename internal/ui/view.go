package ui

import (
	"fmt"
	"strings"

	"gokana/internal/model"

	"github.com/charmbracelet/lipgloss"
)

func View(m *model.Model) string {
	if m.Quitting {
		points := m.GetPoints()
		finalScore := fmt.Sprintf("\nFinal Score: %d points (%d correct)\n", points, m.Correct)
		return finalScore
	}

	if m.GameOver {
		points := m.GetPoints()
		gameOverMsg := fmt.Sprintf("\n💀 GAME OVER 💀\n\nFinal Score: %d points (%d correct)\n\nPress ESC or Ctrl+C to quit\n", points, m.Correct)
		return gameOverMsg
	}

	var s strings.Builder

	s.WriteString(TitleStyle.Render("🗾 Hiragana Quiz"))
	s.WriteString("\n\n")

	livesText := ""
	for i := 0; i < m.Lives; i++ {
		livesText += "❤️  "
	}
	if m.Lives == 0 {
		livesText = "💀 "
	}

	level := m.GetLevel()
	levelText := fmt.Sprintf("🎯 Level %d", level)

	points := m.GetPoints()
	scoreText := fmt.Sprintf("⭐ %dpt", points)

	statsLine := livesText + "  " + levelText + "  " + scoreText
	centeredStats := lipgloss.NewStyle().
		Width(60).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("111")).
		Render(statsLine)

	s.WriteString(centeredStats)
	s.WriteString("\n\n")

	var playArea strings.Builder
	for row := 0; row < m.MaxFallHeight; row++ {
		// Create a map of positions to rendered kana strings
		positionedKanas := make(map[int]string)
		maxPos := 0

		for _, fk := range m.FallingKanas {
			if fk.FallPosition == row {
				var kana string
				if fk.ShowingCorrect {
					correctKanaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
					kana = correctKanaStyle.Render(fk.Kana.Character)
				} else {
					kana = HiraganaStyle.Render(fk.Kana.Character)
				}
				positionedKanas[fk.HorizontalPos] = kana
				if fk.HorizontalPos > maxPos {
					maxPos = fk.HorizontalPos
				}
			}
		}

		// Build the line with spaces and kanas at absolute positions
		// Always build to full width to prevent centering shifts
		line := ""
		for pos := 0; pos < m.PlayAreaWidth; pos++ {
			if kana, exists := positionedKanas[pos]; exists {
				line += kana
			} else {
				line += " "
			}
		}

		playArea.WriteString(line)
		if row < m.MaxFallHeight-1 {
			playArea.WriteString("\n")
		}
	}

	s.WriteString(PlayAreaStyle.Render(playArea.String()))
	s.WriteString("\n\n")

	borderColor := lipgloss.Color("240")
	hasShowingCorrect := m.HasShowingCorrect()
	if hasShowingCorrect {
		borderColor = lipgloss.Color("42")
	} else if m.ShowingFeedback && m.FeedbackType == "wrong" {
		borderColor = lipgloss.Color("196")
	}

	inputBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(30).
		Align(lipgloss.Center).
		Padding(0, 1)

	inputDisplay := m.Input
	if m.Input == "" {
		inputDisplay = "_"
	}

	centeredInputBox := lipgloss.NewStyle().
		Width(60).
		Align(lipgloss.Center).
		Render(inputBoxStyle.Render(inputDisplay))

	s.WriteString(centeredInputBox)
	s.WriteString("\n\n")

	helpText := "Press ESC or Ctrl+C to quit"
	footer := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(helpText)
	s.WriteString(footer)

	return s.String()
}
