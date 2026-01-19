package game

import (
	"strings"
	"time"

	"gokana/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time
type correctDelayMsg time.Time
type feedbackDelayMsg time.Time

const refreshRate = time.Millisecond * 100

func tick() tea.Cmd {
	return tea.Tick(refreshRate, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func correctDelay() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return correctDelayMsg(t)
	})
}

func feedbackDelay() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return feedbackDelayMsg(t)
	})
}

func Init() tea.Cmd {
	return tick()
}

func Update(m *model.Model, msg tea.Msg) (*model.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case correctDelayMsg:
		// Remove all kanas marked as showingCorrect
		newFalling := []model.FallingKana{}
		for _, fk := range m.FallingKanas {
			if !fk.ShowingCorrect {
				newFalling = append(newFalling, fk)
			}
		}
		m.FallingKanas = newFalling

		// Spawn new kanas to maintain level count
		level := m.GetLevel()
		for len(m.FallingKanas) < level {
			m.FallingKanas = append(m.FallingKanas, SpawnKana(m.PlayAreaWidth))
		}

		m.Input = ""
		m.TimeAccumulated = 0
		return m, nil

	case feedbackDelayMsg:
		m.ShowingFeedback = false
		m.Feedback = ""
		m.FeedbackType = ""
		return m, nil

	case tickMsg:
		if m.Quitting || m.GameOver {
			return m, nil
		}

		m.TimeAccumulated += refreshRate
		if m.TimeAccumulated >= m.FallSpeed {
			m.TimeAccumulated -= m.FallSpeed

			// Move all falling kanas
			newFalling := []model.FallingKana{}
			var cmd tea.Cmd = nil

			for _, fk := range m.FallingKanas {
				if fk.ShowingCorrect {
					// Don't move kanas that are showing correct
					newFalling = append(newFalling, fk)
					continue
				}

				fk.FallPosition++
				if fk.FallPosition >= m.MaxFallHeight {
					// Kana reached bottom - lose life and spawn new one
					m.Lives--
					m.Total++
					if m.Lives <= 0 {
						m.GameOver = true
						m.FeedbackType = "wrong"
						m.ShowingFeedback = true
						cmd = feedbackDelay()
					} else {
						m.FeedbackType = "wrong"
						m.ShowingFeedback = true
						if cmd == nil {
							cmd = feedbackDelay()
						}
					}
					m.Input = ""
					// Spawn replacement
					newFalling = append(newFalling, SpawnKana(m.PlayAreaWidth))
				} else {
					newFalling = append(newFalling, fk)
				}
			}

			m.FallingKanas = newFalling
			if cmd != nil {
				return m, tea.Batch(tick(), cmd)
			}
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit

		case tea.KeyBackspace:
			if m.GameOver {
				return m, nil
			}

			if m.ShowingFeedback {
				return m, nil
			}

			// Check if any kana is showing correct
			hasShowingCorrect := m.HasShowingCorrect()

			if hasShowingCorrect {
				m.Input = ""
				// Clear all showingCorrect flags
				for i := range m.FallingKanas {
					m.FallingKanas[i].ShowingCorrect = false
				}
				return m, nil
			}

			if len(m.Input) > 0 {
				m.Input = m.Input[:len(m.Input)-1]
				m.Feedback = ""
			}

		case tea.KeyRunes:
			if m.GameOver {
				return m, nil
			}

			if m.ShowingFeedback {
				return m, nil
			}

			// Check if any kana is showing correct - clear input and reset
			hasShowingCorrect := m.HasShowingCorrect()

			if hasShowingCorrect {
				m.Input = ""
				// Clear all showingCorrect flags
				for i := range m.FallingKanas {
					m.FallingKanas[i].ShowingCorrect = false
				}
			}

			m.Input += string(msg.Runes)
			answer := strings.TrimSpace(strings.ToLower(m.Input))

			// Check if answer matches any falling kana
			matchedIndex := -1
			maxLen := 0
			for i, fk := range m.FallingKanas {
				if answer == fk.Kana.Romaji {
					matchedIndex = i
					break
				}
				// Track max length for wrong answer detection
				if len(fk.Kana.Romaji) > maxLen {
					maxLen = len(fk.Kana.Romaji)
				}
			}

			if matchedIndex != -1 {
				// Correct match!
				m.Total++
				m.Correct++
				m.FeedbackType = "correct"
				m.FallingKanas[matchedIndex].ShowingCorrect = true
				m.TimeAccumulated = 0

				if m.Correct%20 == 0 {
					m.FallSpeed = time.Duration(float64(m.FallSpeed) * 0.85)
					if m.FallSpeed < time.Millisecond*100 {
						m.FallSpeed = time.Millisecond * 100
					}

					// Level up - add more kanas
					level := m.GetLevel()
					for len(m.FallingKanas) < level {
						m.FallingKanas = append(m.FallingKanas, SpawnKana(m.PlayAreaWidth))
					}
				}

				return m, correctDelay()
			} else if len(answer) >= maxLen {
				// Wrong answer - input is long enough
				m.FeedbackType = "wrong"
				m.ShowingFeedback = true
				m.Input = ""
				return m, feedbackDelay()
			} else {
				m.FeedbackType = ""
			}
		}
	}

	return m, nil
}
