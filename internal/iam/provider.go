// Package iam defines the core domain models and interfaces for the application.
package iam

// Provider defines the interface for a cloud provider, abstracting the
// underlying implementation for fetching cloud information.
type Provider interface {
	// GetProjectID returns the currently configured GCP Project ID.
	GetProjectID() string

	// GetAuthenticatedUser returns the email of the currently authenticated user.
	GetAuthenticatedUser() string
}