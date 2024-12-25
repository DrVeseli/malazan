package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var worldMap WorldMap

type tickMsg time.Time

type model struct{}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func init() {
	player, err := LoadPlayer()
	if err != nil {
		player = CreatePlayer()
		SavePlayer(player)
		player, _ = LoadPlayer()
	}

	var err2 error
	worldMap, err2 = LoadWorldMap()
	if err2 != nil {
		worldMap = CreateWorldMap()
		SaveWorldMap(worldMap)
		worldMap, _ = LoadWorldMap()
	}

	SetPlayerLocation(player, &worldMap)
	SaveWorldMap(worldMap)
	go spawnScout()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "Q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
