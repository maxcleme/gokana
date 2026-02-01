package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	KanaStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Bold(true)

	PlayAreaStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Width(60).
		Height(12)

	InputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	CorrectStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("42")).
		Bold(true)

	WrongStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true)

	ScoreStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("111")).
		MarginTop(1)

	HelpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		MarginTop(2)
)
