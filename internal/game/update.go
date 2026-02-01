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
	return nil
}

func Update(m *model.Model, msg tea.Msg) (*model.Model, tea.Cmd) {
	switch m.State {
	case model.StateMenu:
		return updateMenu(m, msg)
	case model.StatePlaying:
		return updatePlaying(m, msg)
	case model.StateGameOver:
		m.Quitting = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func updateMenu(m *model.Model, msg tea.Msg) (*model.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit
		case tea.KeyUp, tea.KeyShiftTab:
			switch m.MenuSection {
			case model.MenuSectionKana:
				m.MenuCursor--
				if m.MenuCursor < 0 {
					m.MenuCursor = 2
				}
			case model.MenuSectionDakuten:
				m.DakutenEnabled = !m.DakutenEnabled
			case model.MenuSectionLevel:
				m.StartLevel++
				if m.StartLevel > 10 {
					m.StartLevel = 1
				}
			case model.MenuSectionLives:
				m.StartLives++
				if m.StartLives > 10 {
					m.StartLives = 1
				}
			}
		case tea.KeyDown, tea.KeyTab:
			switch m.MenuSection {
			case model.MenuSectionKana:
				m.MenuCursor++
				if m.MenuCursor > 2 {
					m.MenuCursor = 0
				}
			case model.MenuSectionDakuten:
				m.DakutenEnabled = !m.DakutenEnabled
			case model.MenuSectionLevel:
				m.StartLevel--
				if m.StartLevel < 1 {
					m.StartLevel = 10
				}
			case model.MenuSectionLives:
				m.StartLives--
				if m.StartLives < 1 {
					m.StartLives = 10
				}
			}
		case tea.KeyLeft:
			m.MenuSection--
			if m.MenuSection < model.MenuSectionKana {
				m.MenuSection = model.MenuSectionStart
			}
		case tea.KeyRight:
			m.MenuSection++
			if m.MenuSection > model.MenuSectionStart {
				m.MenuSection = model.MenuSectionKana
			}
		case tea.KeyEnter, tea.KeySpace:
			if m.MenuSection == model.MenuSectionKana {
				m.SelectedKana = model.KanaType(m.MenuCursor)
				m.MenuSection = model.MenuSectionDakuten
			} else if m.MenuSection == model.MenuSectionStart {
				StartGame(m)
				return m, tick()
			} else {
				m.MenuSection++
			}
		}
	}
	return m, nil
}

func updatePlaying(m *model.Model, msg tea.Msg) (*model.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case correctDelayMsg:
		newFalling := []model.FallingKana{}
		for _, fk := range m.FallingKanas {
			if !fk.ShowingCorrect {
				newFalling = append(newFalling, fk)
			}
		}
		m.FallingKanas = newFalling

		level := m.GetLevel()
		for len(m.FallingKanas) < level {
			m.FallingKanas = append(m.FallingKanas, SpawnKana(m))
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

			newFalling := []model.FallingKana{}
			var cmd tea.Cmd = nil

			for _, fk := range m.FallingKanas {
				if fk.ShowingCorrect {
					newFalling = append(newFalling, fk)
					continue
				}

				fk.FallPosition++
				if fk.FallPosition >= m.MaxFallHeight {
					m.Lives--
					m.Total++
					if m.Lives <= 0 {
						m.GameOver = true
						m.State = model.StateGameOver
						m.Quitting = true
						return m, tea.Quit
					} else {
						m.FeedbackType = "wrong"
						m.ShowingFeedback = true
						if cmd == nil {
							cmd = feedbackDelay()
						}
					}
					m.Input = ""
					newFalling = append(newFalling, SpawnKana(m))
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
			hasShowingCorrect := m.HasShowingCorrect()
			if hasShowingCorrect {
				m.Input = ""
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
			hasShowingCorrect := m.HasShowingCorrect()
			if hasShowingCorrect {
				m.Input = ""
				for i := range m.FallingKanas {
					m.FallingKanas[i].ShowingCorrect = false
				}
			}

			m.Input += string(msg.Runes)
			answer := strings.TrimSpace(strings.ToLower(m.Input))

			matchedIndex := -1
			isValidPrefix := false
			for i, fk := range m.FallingKanas {
				if answer == fk.Kana.Romaji {
					matchedIndex = i
					isValidPrefix = true
					break
				}
				if strings.HasPrefix(fk.Kana.Romaji, answer) {
					isValidPrefix = true
				}
			}

			if matchedIndex != -1 {
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
					level := m.GetLevel()
					for len(m.FallingKanas) < level {
						m.FallingKanas = append(m.FallingKanas, SpawnKana(m))
					}
				}
				return m, correctDelay()
			} else if !isValidPrefix {
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
