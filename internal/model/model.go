package model

import "time"

// Model represents the game state
type Model struct {
	FallingKanas    []FallingKana
	Input           string
	Feedback        string
	FeedbackType    string
	Correct         int
	Total           int
	Lives           int
	Quitting        bool
	GameOver        bool
	MaxFallHeight   int
	PlayAreaWidth   int
	ShowingFeedback bool
	FallSpeed       time.Duration
	TimeAccumulated time.Duration
}

// GetLevel returns the current level based on correct answers
func (m *Model) GetLevel() int {
	return (m.Correct / 20) + 1
}

// GetPoints returns the current points
func (m *Model) GetPoints() int {
	return m.Correct * 100
}

// HasShowingCorrect checks if any kana is showing as correct
func (m *Model) HasShowingCorrect() bool {
	for _, fk := range m.FallingKanas {
		if fk.ShowingCorrect {
			return true
		}
	}
	return false
}
