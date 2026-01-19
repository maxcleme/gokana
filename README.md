# 🗾 Gokana

A gamified CLI hiragana quiz application inspired by [Tofugu's Kana Quiz](https://kana-quiz.tofugu.com/).

## Motivation

This entire repository was **vibecoded** - built through conversational AI-assisted development. The goal was to create a Tofugu-like learning experience but with additional gamification elements to increase engagement while learning hiragana. Instead of just answering questions one by one, you get falling characters, lives, levels, and progressive difficulty to make practice more engaging.

## Features

- 🎮 **Falling hiragana mechanics** - Characters fall from top to bottom, type the romaji before they hit the ground
- ❤️ **Lives system** - Start with 5 lives, lose one when a kana reaches the bottom
- 📈 **Progressive difficulty** - Speed increases and more hiragana appear as you level up
- 🎯 **Level-based gameplay** - Every 20 correct answers = new level with faster speed and more falling kana
- ⭐ **Points system** - Earn 100 points per correct answer
- 🎨 **Clean TUI** - Built with Bubble Tea and Lipgloss for a polished terminal experience
- 🚀 **Start at any level** - Skip the early levels if you're already proficient

## Installation

```bash
go build -o gokana
```

## Usage

Start at level 1 (default):
```bash
./gokana
```

Start at a specific level:
```bash
./gokana 5
```

### Controls

- **Type the romaji** for any falling hiragana and press the corresponding keys
- **Backspace** to correct mistakes
- **ESC or Ctrl+C** to quit

## How It Works

- **Level 1**: 1 falling hiragana, 700ms fall speed
- **Level 2+**: Number of simultaneous hiragana = level number
- **Speed**: Increases by 15% every 20 correct answers (minimum 100ms)
- **Lives**: Lose one when hiragana reaches bottom, game over at 0 lives

## Project Structure

```
gokana/
├── main.go                    # Entry point
├── internal/
│   ├── model/
│   │   ├── kana.go           # Kana types and hiragana data
│   │   └── model.go          # Game state model
│   ├── game/
│   │   ├── game.go           # Game initialization and spawning
│   │   └── update.go         # Game logic and state updates
│   └── ui/
│       ├── styles.go         # Lipgloss styling definitions
│       └── view.go           # View rendering logic
```

## Technical Details

- **Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (TUI framework)
- **Styling**: [Lipgloss](https://github.com/charmbracelet/lipgloss)
- **Architecture**: Model-View-Update (MVU) pattern
- **Rendering**: Time-based animation with 100ms refresh rate
- **Positioning**: Absolute positioning using coordinate maps to prevent UI shifts

## Supported Hiragana

All 46 main hiragana characters:
```
あいうえお (a i u e o)
かきくけこ (ka ki ku ke ko)
さしすせそ (sa shi su se so)
たちつてと (ta chi tsu te to)
なにぬねの (na ni nu ne no)
はひふへほ (ha hi fu he ho)
まみむめも (ma mi mu me mo)
やゆよ (ya yu yo)
らりるれろ (ra ri ru re ro)
わをん (wa wo n)
```

## License

MIT
