package ui

import (
	"fmt"
	"sort"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/tidjee-dev/doit/internal/config"
)

func PrintTaskHeader(category, name, desc string) {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		CategoryStyle.Render(category),
		TaskNameStyle.Render(name),
		TaskDescStyle.Render(desc),
	)

	fmt.Println(TaskBoxStyle.Render(content))
}

func PrintHeading(title string) {
	fmt.Println(HeadingStyle.Render(title))
	fmt.Println()
}

func PrintSection(title string) {
	fmt.Println(SectionStyle.Render(title))
	fmt.Println(SubtleStyle.Render("────────────────────────"))
}

func PrintHelp(cfg config.Config) {
	PrintHeading("doit")

	fmt.Println(SubtleStyle.Render(
		fmt.Sprintf("Task runner for project: %s\n", cfg.App.Name),
	))

	PrintSection("Usage")
	fmt.Printf("  doit <task>\n\n")

	PrintSection("Available tasks")

	categories := map[string][]string{}
	for name, task := range cfg.Tasks {
		categories[task.Category] = append(categories[task.Category], name)
	}

	for _, category := range sortedKeys(categories) {
		PrintSection(category)

		for _, name := range categories[category] {
			task := cfg.Tasks[name]
			fmt.Printf("  %-15s %s\n", name, task.Description)
		}
		fmt.Println()
	}
}

func PrintCommand(cmd string) {
	fmt.Println(" ", CommandStyle.Render("✓"), cmd)
}

func PrintTaskFooter(d time.Duration) {
	fmt.Println(
		FooterStyle.Render(
			fmt.Sprintf("✓ Completed in %s", d.Round(time.Millisecond)),
		),
	)
	fmt.Println()
}

func PrintSummary(tasks int, d time.Duration) {
	fmt.Println(
		SummaryStyle.Render(
			fmt.Sprintf("✓ %d tasks executed in %s",
				tasks,
				d.Round(time.Millisecond),
			),
		),
	)
}

func sortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
