#!/bin/bash

# --- CI Pipeline Tailing Script ---
# Author: Harry O'Malley
#
# This script monitors the latest GitHub Actions workflow run for the current branch.
# It waits for the run to start, streams the live logs, and upon failure,
# prints the logs for the failed jobs.

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
COMMIT_SHA=$(git rev-parse HEAD)
if [ -z "$COMMIT_SHA" ]; then
    echo "Error: Could not determine the current commit SHA." >&2
    exit 1
fi

RUN_ID=""
echo "- Searching for workflow run for commit ${COMMIT_SHA:0:7}..."

# Loop for a maximum of 2 minutes (12 attempts * 10s sleep) waiting for the run to be created.
for i in {1..12}; do
    RUN_ID=$(gh run list --commit "$COMMIT_SHA" --limit 1 --json databaseId -q 'map(.databaseId) | .[0]')
    if [ -n "$RUN_ID" ] && [ "$RUN_ID" != "null" ]; then
        echo "- Found run #$RUN_ID."
        break
    fi
    echo "- Workflow has not started yet, waiting 10 seconds... (Attempt $i/12)"
    sleep 10
done

if [ -z "$RUN_ID" ] || [ "$RUN_ID" == "null" ]; then
    echo "- Error: Timed out waiting for workflow run to start for commit ${COMMIT_SHA:0:7}." >&2
    exit 1
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