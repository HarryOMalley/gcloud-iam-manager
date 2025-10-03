package gcp

import (
	"context"
	"fmt"
	"gcloud-iam-manager/internal/iam"
	"os/exec"
	"strings"

	googleiam "google.golang.org/api/iam/v1"
)

// Client implements the iam.Provider interface for Google Cloud Platform.
type Client struct {
	projectID string
	user      string
	iamService *googleiam.Service
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

	ctx := context.Background()
	iamService, err := googleiam.NewService(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "could not find default credentials") {
			return nil, fmt.Errorf("authentication failed: could not find default credentials. Please run `gcloud auth application-default login` to authenticate")
		}
		return nil, err
	}

	return &Client{
		projectID: projectID,
		user:      user,
		iamService: iamService,
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

// GetRoles returns a list of all predefined IAM roles available in GCP.
func (c *Client) GetRoles() ([]*iam.Role, error) {
	var roles []*iam.Role

	pageToken := ""
	for {
		req := c.iamService.Roles.List().View("FULL").PageToken(pageToken)
		resp, err := req.Do()
		if err != nil {
			return nil, err
		}

		for _, role := range resp.Roles {
			// Filter out roles that are deleted or in early access stage.
			if role.Stage != "DELETED" && role.Stage != "EAP" {
				roles = append(roles, &iam.Role{
					Name:        role.Name,
					Title:       role.Title,
					Description: role.Description,
				})
			}
		}

		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}

	return roles, nil
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
