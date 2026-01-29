package ui

import "github.com/charmbracelet/lipgloss"

var (
	HeadingStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("69")).
			Border(lipgloss.DoubleBorder()).
			Padding(0, 2).
			Align(lipgloss.Center)

	SectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("220"))

	SubtleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	SuccessStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42"))

	WarnStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214"))

	ErrorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))

	InfoStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("81"))

	TaskBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1).
			MarginTop(1)

	TaskTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("69"))

	TaskDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	CommandOKStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	DependencyBoxStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("245")).
				BorderForeground(lipgloss.Color("240"))

	CategoryStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("81"))

	TaskNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("69"))

	CommandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MarginLeft(2)

	SummaryStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")).
			MarginTop(1)
)
