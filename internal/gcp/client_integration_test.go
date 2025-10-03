//go:build integration

package gcp

import (
	"testing"
)

func TestGetRoles(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create client for testing roles: %v", err)
	}

	roles, err := client.GetRoles()
	if err != nil {
		t.Fatalf("GetRoles() returned an unexpected error: %v", err)
	}

	if len(roles) == 0 {
		t.Fatal("GetRoles() returned no roles. Expected a list of predefined roles.")
	}

	// Sanity check for a common, well-known role
	foundViewer := false
	for _, role := range roles {
		if role.Name == "roles/viewer" {
			foundViewer = true
			if role.Title != "Viewer" {
				t.Errorf("Expected roles/viewer title to be 'Viewer', got %q", role.Title)
			}
			break
		}
	}

	if !foundViewer {
		t.Error("Expected to find the predefined 'roles/viewer' role, but it was missing.")
	}
}

func TestNewClient(t *testing.T) {
	// 1. Get expected values directly from the gcloud CLI.
	expectedProjectID := getGcloudOutput(t, "config", "get-value", "project")
	expectedUser := getGcloudOutput(t, "auth", "list", "--filter=status:ACTIVE", "--format=value(account)")

	// 2. Call the function we are testing.
	client, err := NewClient()

	// 3. Assert the results.
	if err != nil {
		t.Fatalf("NewClient() returned an unexpected error: %v", err)
	}

	if client == nil {
		t.Fatal("NewClient() returned a nil client")
	}

	projectID := client.GetProjectID()
	if projectID != expectedProjectID {
		t.Errorf("Expected project ID %q, but got %q", expectedProjectID, projectID)
	}

	user := client.GetAuthenticatedUser()
	if user != expectedUser {
		t.Errorf("Expected user %q, but got %q", expectedUser, user)
	}
}