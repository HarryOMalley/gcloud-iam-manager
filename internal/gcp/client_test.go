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