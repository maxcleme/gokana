package game

import (
	"math/rand"
	"time"

	"gokana/internal/model"
)

func SpawnKana(m *model.Model) model.FallingKana {
	kanaSet := model.GetKanaSet(m.SelectedKana, m.DakutenEnabled)
	return model.FallingKana{
		Kana:           kanaSet[rand.Intn(len(kanaSet))],
		FallPosition:   0,
		HorizontalPos:  rand.Intn(m.PlayAreaWidth),
		ShowingCorrect: false,
	}
}

func InitialModel() *model.Model {
	playWidth := 55
	return &model.Model{
		State:           model.StateMenu,
		SelectedKana:    model.KanaTypeBoth,
		DakutenEnabled:  true,
		MenuCursor:      2,
		MenuSection:     model.MenuSectionStart,
		StartLevel:      1,
		StartLives:      4,
		FallingKanas:    []model.FallingKana{},
		Correct:         0,
		MaxFallHeight:   15,
		PlayAreaWidth:   playWidth,
		FallSpeed:       time.Millisecond * 700,
		TimeAccumulated: 0,
		Lives:           4,
	}
}

func StartGame(m *model.Model) {
	startLevel := m.StartLevel
	if startLevel < 1 {
		startLevel = 1
	}
	if startLevel > 10 {
		startLevel = 10
	}

	// Calculate speed based on level
	speed := time.Millisecond * 700
	for i := 1; i < startLevel; i++ {
		speed = time.Duration(float64(speed) * 0.85)
		if speed < time.Millisecond*100 {
			speed = time.Millisecond * 100
			break
		}
	}

	m.State = model.StatePlaying
	m.FallSpeed = speed
	m.Correct = 0
	m.Total = 0
	m.LevelOffset = startLevel - 1
	m.Lives = m.StartLives
	m.GameOver = false
	m.Quitting = false
	m.Input = ""
	m.Feedback = ""
	m.FeedbackType = ""
	m.ShowingFeedback = false
	m.TimeAccumulated = 0
	m.FallingKanas = []model.FallingKana{}

	// Spawn initial kanas based on level
	for i := 0; i < startLevel; i++ {
		m.FallingKanas = append(m.FallingKanas, SpawnKana(m))
	}
}
