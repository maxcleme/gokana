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
		if m.GameOver {
			return fmt.Sprintf("\n💀 GAME OVER 💀\n\nFinal Score: %d points (%d correct)\n", points, m.Correct)
		}
		return fmt.Sprintf("\nFinal Score: %d points (%d correct)\n", points, m.Correct)
	}

	switch m.State {
	case model.StateMenu:
		return viewMenu(m)
	case model.StatePlaying:
		return viewGame(m)
	default:
		return ""
	}
}

func viewMenu(m *model.Model) string {
	var s strings.Builder

	s.WriteString(TitleStyle.Render("🗾 Gokana"))
	s.WriteString("\n\n")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("111")).
		Italic(true).
		Render("Japanese Kana Quiz Game")
	s.WriteString(subtitle)
	s.WriteString("\n\n")

	sectionStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	activeSectionStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("111"))
	activeValueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	// Kana Selection
	kanaHeader := "Character Set:"
	if m.MenuSection == model.MenuSectionKana {
		s.WriteString(activeSectionStyle.Render("▸ " + kanaHeader))
	} else {
		s.WriteString(sectionStyle.Render("  " + kanaHeader))
	}
	s.WriteString("\n")

	options := []struct {
		name string
		desc string
	}{
		{"Hiragana", "あ い う え お"},
		{"Katakana", "ア イ ウ エ オ"},
		{"Both", "あ ア い イ う ウ"},
	}

	for i, opt := range options {
		cursor := "    "
		var optStyle lipgloss.Style
		if m.MenuSection == model.MenuSectionKana {
			if i == m.MenuCursor {
				cursor = "  ▸ "
				optStyle = activeValueStyle
			} else {
				optStyle = dimStyle
			}
		} else {
			if i == int(m.SelectedKana) {
				cursor = "  ✓ "
				optStyle = valueStyle
			} else {
				optStyle = dimStyle
			}
		}
		line := cursor + optStyle.Render(opt.name) + "  " + dimStyle.Render(opt.desc)
		s.WriteString(line + "\n")
	}
	s.WriteString("\n")

	// Dakuten Selection
	dakutenHeader := "Include Dakuten:"
	if m.MenuSection == model.MenuSectionDakuten {
		s.WriteString(activeSectionStyle.Render("▸ " + dakutenHeader))
		s.WriteString("  ")
		if m.DakutenEnabled {
			s.WriteString(activeValueStyle.Render("< ON >"))
			s.WriteString("  " + dimStyle.Render("が ざ だ ば ぱ"))
		} else {
			s.WriteString(activeValueStyle.Render("< OFF >"))
			s.WriteString("  " + dimStyle.Render("basic kana only"))
		}
	} else {
		s.WriteString(sectionStyle.Render("  " + dakutenHeader))
		s.WriteString("  ")
		if m.DakutenEnabled {
			s.WriteString(valueStyle.Render("ON"))
			s.WriteString("  " + dimStyle.Render("が ざ だ ば ぱ"))
		} else {
			s.WriteString(valueStyle.Render("OFF"))
			s.WriteString("  " + dimStyle.Render("basic kana only"))
		}
	}
	s.WriteString("\n\n")

	// Level Selection
	levelHeader := "Starting Level:"
	if m.MenuSection == model.MenuSectionLevel {
		s.WriteString(activeSectionStyle.Render("▸ " + levelHeader))
		s.WriteString("  ")
		s.WriteString(activeValueStyle.Render(fmt.Sprintf("< %d >", m.StartLevel)))
	} else {
		s.WriteString(sectionStyle.Render("  " + levelHeader))
		s.WriteString("  ")
		s.WriteString(valueStyle.Render(fmt.Sprintf("%d", m.StartLevel)))
	}
	s.WriteString("\n\n")

	// Lives Selection
	livesHeader := "Starting Lives:"
	if m.MenuSection == model.MenuSectionLives {
		s.WriteString(activeSectionStyle.Render("▸ " + livesHeader))
		s.WriteString("  ")
		hearts := ""
		for i := 0; i < m.StartLives; i++ {
			hearts += "❤️ "
		}
		s.WriteString(activeValueStyle.Render(fmt.Sprintf("< %d >", m.StartLives)))
		s.WriteString("  " + hearts)
	} else {
		s.WriteString(sectionStyle.Render("  " + livesHeader))
		s.WriteString("  ")
		hearts := ""
		for i := 0; i < m.StartLives; i++ {
			hearts += "❤️ "
		}
		s.WriteString(valueStyle.Render(fmt.Sprintf("%d", m.StartLives)))
		s.WriteString("  " + hearts)
	}
	s.WriteString("\n\n")

	// Start Button
	if m.MenuSection == model.MenuSectionStart {
		startBtnStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color("205")).
			Padding(0, 2)
		s.WriteString(startBtnStyle.Render("▸ START GAME"))
	} else {
		startBtnStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 2)
		s.WriteString(startBtnStyle.Render("  START GAME"))
	}
	s.WriteString("\n\n")

	helpText := dimStyle.Render("←/→ sections • ↑/↓ adjust • Enter to confirm • ESC to quit")
	s.WriteString(helpText)

	return s.String()
}

func viewGame(m *model.Model) string {
	var s strings.Builder

	title := fmt.Sprintf("🗾 %s Quiz", m.SelectedKana.String())
	s.WriteString(TitleStyle.Render(title))
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
		positionedKanas := make(map[int]string)
		maxPos := 0

		for _, fk := range m.FallingKanas {
			if fk.FallPosition == row {
				var kana string
				if fk.ShowingCorrect {
					correctKanaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
					kana = correctKanaStyle.Render(fk.Kana.Character)
				} else {
					kana = KanaStyle.Render(fk.Kana.Character)
				}
				positionedKanas[fk.HorizontalPos] = kana
				if fk.HorizontalPos > maxPos {
					maxPos = fk.HorizontalPos
				}
			}
		}

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
