package gcp

import (
	"os/exec"
	"strings"
	"testing"
)

// getGcloudOutput is a helper to shell out to the gcloud command.
// This is a simple approach for an initial integration test.
// In a real-world scenario with more complex logic, we would mock this.
func getGcloudOutput(t *testing.T, args ...string) string {
	t.Helper()
	cmd := exec.Command("gcloud", args...)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("gcloud command failed: %v. Ensure you are authenticated with gcloud.", err)
	}
	return strings.TrimSpace(string(out))
}

func TestNewClient(t *testing.T) {
	// 1. Get expected values directly from the gcloud CLI.
	expectedProjectID := getGcloudOutput(t, "config", "get-value", "project")
	expectedUser := getGcloudOutput(t, "auth", "list", "--filter=status:ACTIVE", "--format=value(account)")

	// 2. Call the function we are testing (which doesn't exist yet).
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
