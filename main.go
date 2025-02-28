package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    board string
}

func initModel() model {
    return model {
        board: "board",
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return ""
}

func main() {
    p := tea.NewProgram(initModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
}
