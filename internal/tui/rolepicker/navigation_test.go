package rolepicker

import (
	"testing"

	"gcloud-iam-manager/internal/iam"
	tea "github.com/charmbracelet/bubbletea"
)

func TestNavigationAndSelection(t *testing.T) {
	roles := []*iam.Role{
		{Name: "roles/viewer"},
		{Name: "roles/editor"},
	}
	m := New(roles)

	// Test moving down
	msgDown := tea.KeyMsg{Type: tea.KeyDown}
	updatedModel, _ := m.Update(msgDown)
	m = updatedModel.(Model)

	if m.cursor != 1 {
		t.Fatalf("Cursor should be at 1 after moving down, but got %d", m.cursor)
	}

	// Test selection with spacebar
	msgSelect := tea.KeyMsg{Type: tea.KeySpace}
	updatedModel, _ = m.Update(msgSelect)
	m = updatedModel.(Model)

	// The role at cursor 1 is "roles/editor"
	if !m.IsSelected("roles/editor") { // IsSelected method doesn't exist yet
		t.Error("Expected 'roles/editor' to be selected, but it wasn't")
	}

	// Test deselection
	updatedModel, _ = m.Update(msgSelect)
	m = updatedModel.(Model)

	if m.IsSelected("roles/editor") {
		t.Error("Expected 'roles/editor' to be deselected, but it was still selected")
	}
}
