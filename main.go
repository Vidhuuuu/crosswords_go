package main

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
    "bufio"
    "strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
    board [][]rune
}

type cell struct {
    x, y int
}

func initModel() model {
    file, err := os.Open("./assets/cw.info")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    // defer func() {
    //     file.Close()
    // }()

    scanner := bufio.NewScanner(file)
    blockedCell := make(map[cell]struct{})
    for scanner.Scan() {
        var c cell
        coords := strings.Split(scanner.Text(), ",")
        c.x, _ = strconv.Atoi(coords[0])
        c.y, _ = strconv.Atoi(coords[1])
        blockedCell[c] = struct{}{}
    }

    _board := make([][]rune, 15)
    min, max := 65, 90
    for i := range 15 {
        _board[i] = make([]rune, 15)
        for j := range 15 {
            if _, exists := blockedCell[cell{i, j}]; !exists {
                _board[i][j] = rune(rand.Intn(max - min + 1) + min)
            } else {
                _board[i][j] = '$'
            }
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
            if m.board[i][j] == '$' {
                s.WriteString(boardStyle.Render("┃") + "    ")
            } else {
                s.WriteString(boardStyle.Render("┃  " + string(m.board[i][j]) + " "))
            }
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
