package main

import (
	// ... (ваши импорты)
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ... (ваши глобальные переменные: words, стили)

type testState struct {
	completeWords int16
	currentWord   string
	latters       []bool
	userWords     []string

	// Новые поля для размеров терминала
	width  int
	height int
}

// ... (функция generateWords)

func initialState(wordCount int) testState {
	return testState{
		completeWords: 0,
		currentWord:   "",
		latters:       []bool{false},
		userWords:     []string{"hello"},

		// На старте размеры равны нулю
		width:  0,
		height: 0,
	}
}

func (s testState) Init() tea.Cmd {
	return tea.SetWindowTitle("Typing zen")
}

func (s testState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Добавьте здесь свою логику обработки клавиш
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return s, tea.Quit
		}

	case tea.WindowSizeMsg:
		// Обработка изменения размера окна
		s.width = msg.Width
		s.height = msg.Height
	}

	return s, nil
}

func (s testState) View() string {
	// Используем lipgloss для создания контейнера на весь экран
	textToDisplay := strings.Join(s.userWords, " ")

	// Создаем стиль, который будет заполнять все доступное пространство
	// и центрировать текст
	containerStyle := lipgloss.NewStyle().
		Width(s.width).
		Height(s.height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	// Рендерим текст внутри этого контейнера
	return containerStyle.Render(textToDisplay)
}

func main() {
	wordsCountFlag := flag.Int("w", 0, "enter words count")
	flag.Parse()

	p := tea.NewProgram(initialState(*wordsCountFlag))

	if _, err := p.Run(); err != nil {
		fmt.Println("Ошибка запуска программы: ", err)
		os.Exit(1)
	}
}
