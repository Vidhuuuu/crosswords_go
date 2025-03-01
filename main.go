package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
    board [][]rune
}

func initModel() model {
    _board := make([][]rune, 15)
    for i := range 15 {
        _board[i] = make([]rune, 15)
        for j := range 15 {
            _board[i][j] = '$'
        }
    }
    return model {
        board: _board,
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
    boardStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("0")).
        Background(lipgloss.Color("15"))

    var s strings.Builder
    s.WriteString("\n\n\t")

    s.WriteString(boardStyle.Render("┏━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┳━━━━┓"))
    for i := range 15 {
        if i != 0 {
            s.WriteString("\n\t")
            s.WriteString(boardStyle.Render("┣━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━╋━━━━┫"))
        }
        s.WriteString("\n\t")
        for j := range 15 {
            s.WriteString(boardStyle.Render("┃  " + string(m.board[i][j]) + " "))
        }
        s.WriteString(boardStyle.Render("┃"))
    }
    s.WriteString("\n\t")
    s.WriteString(boardStyle.Render("┗━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┻━━━━┛"))
    return s.String()
}

func main() {
    p := tea.NewProgram(initModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
}
