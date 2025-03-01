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
    posX, posY int
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
    _x, _y := -1, -1
    min, max := 65, 90
    for i := range 15 {
        _board[i] = make([]rune, 15)
        for j := range 15 {
            if _, exists := blockedCell[cell{i, j}]; !exists {
                if(_x == -1 && _y == -1) {
                    _x = i;
                    _y = j;
                }
                _board[i][j] = rune(rand.Intn(max - min + 1) + min)
            } else {
                _board[i][j] = '$'
            }
        }
    }
    return model {
        board: _board,
        posX: _x,
        posY: _y,
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
        case "up":
            if(m.posX > 0) {
                m.posX--;
            }
        case "down":
            if(m.posX < 14) {
                m.posX++;
            }
        case "left":
            if(m.posY > 0) {
                m.posY--;
            }
        case "right":
            if(m.posY < 14) {
                m.posY++;
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    boardStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("0")).
        Background(lipgloss.Color("15"))

    activeStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("0")).
        Background(lipgloss.Color("105"))

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
            } else if (i == m.posX && j == m.posY) {
                s.WriteString(
                    boardStyle.Render("┃") +
                    activeStyle.Render("  " + string(m.board[i][j]) + " "))
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
