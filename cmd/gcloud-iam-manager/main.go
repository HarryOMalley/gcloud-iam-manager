package main

import (
	"fmt"
	"os"

	"gcloud-iam-manager/internal/gcp"
	"gcloud-iam-manager/internal/iam"
	tea "github.com/charmbracelet/bubbletea"
)

// --- Model ---

type model struct {
	projectID string
	user      string
	choices   []string // The menu choices
	cursor    int      // Which choice is currently selected
	err       error
}

func initialModel() model {
	return model{
		choices: []string{"Create Service Account"},
	}
}

// --- Messages ---

type gcpInfoLoadedMsg struct {
	Provider iam.Provider
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

// --- Commands ---

func fetchGCPInfo() tea.Msg {
	client, err := gcp.NewClient()
	if err != nil {
		return errMsg{err}
	}
	return gcpInfoLoadedMsg{Provider: client}
}

// --- TUI Methods ---

func (m model) Init() tea.Cmd {
	return fetchGCPInfo
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}

	case gcpInfoLoadedMsg:
		m.projectID = msg.Provider.GetProjectID()
		m.user = msg.Provider.GetAuthenticatedUser()
		return m, nil

	case errMsg:
		m.err = msg
		return m, tea.Quit // Quit on error for now
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.projectID == "" || m.user == "" {
		return "Loading GCP information..."
	}

	// Header
	header := fmt.Sprintf("GCP IAM Manager\n\nProject: %s\nUser:    %s\n", m.projectID, m.user)

	// Menu
	menu := ""
	for i, choice := range m.choices {
		cursor := " " // not selected
		if m.cursor == i {
			cursor = ">" // selected
		}
		menu += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// Footer
	footer := "\n(Press 'q' to quit)"

	return header + "\n" + menu + footer
}

// --- Main ---

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}
