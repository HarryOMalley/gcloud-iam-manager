# Google Cloud IAM Manager

## Overview

`gcloud-iam-manager` is a modern, terminal-based UI for managing Google Cloud IAM resources. It provides a fast and intuitive way to handle common tasks like creating service accounts, managing user roles, and auditing permissions, without leaving the command line.

The tool is designed to be highly extensible, allowing for the gradual addition of more complex IAM management features.

For details on how to contribute, please see our [Contributing Guidelines](CONTRIBUTING.md).

## Features

- Load and display the current GCP project and authenticated user.
- List all users and service accounts with access to the project.
- An interactive, hierarchical role-picker for creating new service accounts.
- Option to generate and save a JSON key for a new service account.

## Roadmap

1. **Initial Setup**: Configure the project structure, CI, and development guidelines.
2. **Read Project Information**: Load and display the current GCP project and user.
3. **Create Service Account**: Implement the workflow to create a new service account with selected roles.
4. **List Existing Users**: Display a list of existing users and service accounts with their roles.
5. **Manage Roles**: Add functionality to grant and revoke roles from existing principals.
