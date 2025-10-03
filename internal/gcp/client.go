package gcp

import (
	"gcloud-iam-manager/internal/iam"
	"os/exec"
	"strings"
)

// Client implements the iam.Provider interface for Google Cloud Platform.
type Client struct {
	projectID string
	user      string
}

// NewClient creates a new GCP client by shelling out to the `gcloud` command.
// It fetches the active project ID and authenticated user.
func NewClient() (*Client, error) {
	projectID, err := execGcloud("config", "get-value", "project")
	if err != nil {
		return nil, err
	}

	user, err := execGcloud("auth", "list", "--filter=status:ACTIVE", "--format=value(account)")
	if err != nil {
		return nil, err
	}

	return &Client{
		projectID: projectID,
		user:      user,
	}, nil
}

// GetProjectID returns the currently configured GCP Project ID.
func (c *Client) GetProjectID() string {
	return c.projectID
}

// GetAuthenticatedUser returns the email of the currently authenticated user.
func (c *Client) GetAuthenticatedUser() string {
	return c.user
}

// execGcloud is a helper function to execute gcloud commands.
func execGcloud(args ...string) (string, error) {
	cmd := exec.Command("gcloud", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// Ensure Client satisfies the iam.Provider interface.
var _ iam.Provider = (*Client)(nil)
