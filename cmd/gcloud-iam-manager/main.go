package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// model defines the state of our application.
type model struct{}

// Init is the first function that will be called. It returns a command.
func (m model) Init() tea.Cmd {
	// We don't need to do anything on startup, so we return nil.
	return nil
}

// Update is called when a message is received. It's where we handle user input.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// tea.KeyMsg is a message sent when the user presses a key.
	case tea.KeyMsg:
		// If the user presses 'q' or 'ctrl+c', we'll quit.
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime.
	return m, nil
}

// View renders the UI. It's called after every Update.
func (m model) View() string {
	return "Hello, World! Welcome to gcloud-iam-manager.\nPress 'q' to quit."
}

func main() {
	// Create a new Bubble Tea program with our initial model.
	p := tea.NewProgram(model{})

	// Run the program.
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}

