package main

import "github.com/charmbracelet/lipgloss"

var (
	blackStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#000000")).SetString("  ")
	greenStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#00FF00")).SetString("  ")
	blueStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#0000FF")).SetString("  ")
	grayStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#808080")).SetString("  ")
	yellowStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FFFF00")).SetString(" O ")
)

func (m model) View() string {
	var s string
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			index := i*100 + j
			switch worldMap[index] {
			case 0:
				s += blackStyle.String()
			case 1:
				s += greenStyle.String()
			case 2:
				s += blueStyle.String()
			case 3:
				s += grayStyle.String()
			case 4:
				s += yellowStyle.String()
			}
		}
		s += "\n"
	}
	return s + "\nPress Q to quit\n"
}
