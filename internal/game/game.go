package game

import (
	"math/rand"
	"time"

	"gokana/internal/model"
)

func SpawnKana(playWidth int) model.FallingKana {
	return model.FallingKana{
		Kana:           model.MainHiragana[rand.Intn(len(model.MainHiragana))],
		FallPosition:   0,
		HorizontalPos:  rand.Intn(playWidth),
		ShowingCorrect: false,
	}
}

func InitialModel(startLevel int) *model.Model {
	playWidth := 55
	if startLevel < 1 {
		startLevel = 1
	}

	// Spawn hiragana based on starting level
	initialKanas := []model.FallingKana{}
	for i := 0; i < startLevel; i++ {
		initialKanas = append(initialKanas, SpawnKana(playWidth))
	}

	// Calculate speed based on level
	// Level 1 = 700ms, each level is 85% of previous
	speed := time.Millisecond * 700
	for i := 1; i < startLevel; i++ {
		speed = time.Duration(float64(speed) * 0.85)
		if speed < time.Millisecond*100 {
			speed = time.Millisecond * 100
			break
		}
	}

	// Set correct count to match the level
	// Level = (correct / 20) + 1, so correct = (level - 1) * 20
	correctCount := (startLevel - 1) * 20
	return &model.Model{
		FallingKanas:    initialKanas,
		Correct:         correctCount,
		MaxFallHeight:   15,
		PlayAreaWidth:   playWidth,
		FallSpeed:       speed,
		TimeAccumulated: 0,
		Lives:           5,
	}
}
