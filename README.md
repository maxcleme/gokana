# 🗾 Gokana

A gamified CLI Japanese kana quiz application inspired by [Tofugu's Kana Quiz](https://kana-quiz.tofugu.com/).

## Motivation

This entire repository was **vibecoded** - built through conversational AI-assisted development. The goal was to create a Tofugu-like learning experience but with additional gamification elements to increase engagement while learning Japanese kana. Instead of just answering questions one by one, you get falling characters, lives, levels, and progressive difficulty to make practice more engaging.

## Features

- 🎮 **Falling kana mechanics** - Characters fall from top to bottom, type the romaji before they hit the ground
- 🔤 **Full kana support** - Practice hiragana, katakana, or both simultaneously
- ゛ **Dakuten & handakuten** - Optional voiced and semi-voiced consonants (が, ぱ, etc.)
- ❤️ **Lives system** - Start with 4 lives (configurable 1-10), lose one when a kana reaches the bottom
- 📈 **Progressive difficulty** - Speed increases and more kana appear as you level up
- 🎯 **Level-based gameplay** - Every 20 correct answers = new level with faster speed and more falling kana
- ⭐ **Points system** - Earn 100 points per correct answer
- 🎨 **Clean TUI** - Built with Bubble Tea and Lipgloss for a polished terminal experience
- 📋 **Interactive menu** - Configure kana type, dakuten, starting level, and lives before playing

## Installation

```bash
go build -o gokana
```

## Usage

```bash
./gokana
```

The game starts with an interactive menu where you can configure:
- **Character Set**: Hiragana, Katakana, or Both
- **Dakuten**: Enable/disable voiced consonants (が, ざ, だ, ば, ぱ, etc.)
- **Starting Level**: 1-10
- **Starting Lives**: 1-10

### Menu Controls

- **←/→** Navigate between sections
- **↑/↓** Adjust values within a section
- **Enter/Space** Confirm selection and move to next section
- **ESC or Ctrl+C** Quit

### Game Controls

- **Type the romaji** for any falling kana
- **Backspace** to correct mistakes
- **ESC or Ctrl+C** to quit

## How It Works

- **Level 1**: 1 falling kana, 700ms fall speed
- **Level 2+**: Number of simultaneous kana = level number
- **Speed**: Increases by 15% every 20 correct answers (minimum 100ms)
- **Lives**: Lose one when kana reaches bottom, game over at 0 lives

## Project Structure

```
gokana/
├── main.go                    # Entry point
├── internal/
│   ├── model/
│   │   ├── kana.go           # Kana types and character data
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

## Supported Characters

### Hiragana (46 main + 25 with dakuten/handakuten)

**Main hiragana:**
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

**Dakuten (voiced):**
```
がぎぐげご (ga gi gu ge go)
ざじずぜぞ (za ji zu ze zo)
だぢづでど (da di du de do)
ばびぶべぼ (ba bi bu be bo)
```

**Handakuten (semi-voiced):**
```
ぱぴぷぺぽ (pa pi pu pe po)
```

### Katakana (46 main + 25 with dakuten/handakuten)

**Main katakana:**
```
アイウエオ (a i u e o)
カキクケコ (ka ki ku ke ko)
サシスセソ (sa shi su se so)
タチツテト (ta chi tsu te to)
ナニヌネノ (na ni nu ne no)
ハヒフヘホ (ha hi fu he ho)
マミムメモ (ma mi mu me mo)
ヤユヨ (ya yu yo)
ラリルレロ (ra ri ru re ro)
ワヲン (wa wo n)
```

**Dakuten (voiced):**
```
ガギグゲゴ (ga gi gu ge go)
ザジズゼゾ (za ji zu ze zo)
ダヂヅデド (da di du de do)
バビブベボ (ba bi bu be bo)
```

**Handakuten (semi-voiced):**
```
パピプペポ (pa pi pu pe po)
```

## License

MIT
