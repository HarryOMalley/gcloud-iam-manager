#!/bin/bash

# --- CI Pipeline Tailing Script ---
# Author: Harry O'Malley
#
# This script monitors the latest GitHub Actions workflow run for the current branch.
# It streams the live logs and, upon failure, prints the logs for the failed jobs.

# --- Help Function ---
show_help() {
    echo "Usage: ./ci-tail.sh"
    echo "Monitors the latest CI run for the current git branch."
}

if [[ "$1" == "--help" ]]; then
    show_help
    exit 0
fi

# --- Prerequisite Check ---
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI (gh) is not installed. Please install it to continue." >&2
    exit 1
fi

# --- Main Logic ---
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ -z "$CURRENT_BRANCH" ]; then
    echo "Error: Could not determine the current git branch." >&2
    exit 1
fi

echo "- Fetching the latest workflow run for branch '$CURRENT_BRANCH'..."

# Get the ID of the most recent run for the current branch.
RUN_ID=$(gh run list --branch "$CURRENT_BRANCH" --limit 1 --json databaseId -q 'map(.databaseId) | .[0]')

if [ -z "$RUN_ID" ] || [ "$RUN_ID" == "null" ]; then
    echo "- No workflow runs found for this branch yet."
    exit 0
fi

echo "- Tailing run #$RUN_ID. Press Ctrl+C to exit."

# Watch the run and stream logs. Exit with an error code if the run fails.
if ! gh run watch "$RUN_ID" --exit-status; then
    echo
    echo "--- 🔴 Run #$RUN_ID Failed ---"
    echo "- Fetching logs for failed jobs..."
    echo
    gh run view "$RUN_ID" --log-failed
    exit 1
fi

echo
echo "--- ✅ Run #$RUN_ID Completed Successfully ---"
