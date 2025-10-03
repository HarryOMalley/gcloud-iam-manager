package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)
// mockProvider is a mock implementation of the iam.Provider for testing.
type mockProvider struct {
	projectID string
	user      string
}

func (m *mockProvider) GetProjectID() string {
	return m.projectID
}

func (m *mockProvider) GetAuthenticatedUser() string {
	return m.user
}

func TestGcpInfoLoaded(t *testing.T) {
	// 1. Define the initial model and the mock data.
	m := initialModel() // This function doesn't exist yet.
	mockInfo := &mockProvider{
		projectID: "test-project",
		user:      "test-user@example.com",
	}

	// 2. Create the message that our Update function will receive.
	// This message and its struct don't exist yet.
	msg := gcpInfoLoadedMsg{Provider: mockInfo}

	// 3. Send the message to the Update function.
	updatedModel, _ := m.Update(msg)

	// 4. Assert that the model's state has been updated correctly.
	// These fields don't exist on the model yet.
	finalModel := updatedModel.(model)
	if finalModel.projectID != "test-project" {
		t.Errorf("Expected projectID to be 'test-project', got %q", finalModel.projectID)
	}
	if finalModel.user != "test-user@example.com" {
		t.Errorf("Expected user to be 'test-user@example.com', got %q", finalModel.user)
	}
}

func TestMenuNavigation(t *testing.T) {
	m := initialModel()
	m.choices = []string{"one", "two", "three"} // The 'choices' field doesn't exist yet

	// Test moving down
	msg := tea.KeyMsg{Type: tea.KeyDown}
	updatedModel, _ := m.Update(msg)
	finalModel := updatedModel.(model)

	if finalModel.cursor != 1 { // The 'cursor' field doesn't exist yet
		t.Errorf("Expected cursor to be 1 after moving down, got %d", finalModel.cursor)
	}

	// Test moving up
	msg = tea.KeyMsg{Type: tea.KeyUp}
	updatedModel, _ = finalModel.Update(msg)
	finalModel = updatedModel.(model)

	if finalModel.cursor != 0 {
		t.Errorf("Expected cursor to be 0 after moving up, got %d", finalModel.cursor)
	}
}