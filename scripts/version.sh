#!/bin/bash

# --- Version Calculation Script ---
# Author: Harry O'Malley
#
# This script calculates and displays version numbers based on git history
# and conventional commits using the git-cliff tool.

# --- Help Function ---
show_help() {
    echo "Usage: ./version.sh [FLAG]"
    echo "Calculates and displays project version numbers."
    echo
    echo "Flags:"
    echo "  --current    Print the raw current version number (latest git tag)."
    echo "  --next       Print the raw next semantic version number based on conventional commits."
    echo "  --feature    Print the raw version number for a feature branch, including a branch-specific suffix."
    echo "  --help       Show this help message."
}

# --- Argument Parsing ---
if [[ "$1" == "--help" ]]; then
    show_help
    exit 0
fi

# --- Prerequisite Check ---
if ! command -v git-cliff &> /dev/null; then
    echo "Error: git-cliff is not installed. Please install it to continue." >&2
    exit 1
fi

# --- Calculations ---
CURRENT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "0.0.0")
NEXT_VERSION=$(git-cliff --bumped-version 2>/dev/null || echo "0.1.0")

# Create a feature-branch-safe suffix
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
if [[ "$BRANCH_NAME" != "main" && "$BRANCH_NAME" != "master" ]]; then
    # Sanitize branch name for use in a URL and version string
    REF_SLUG=$(echo "$BRANCH_NAME" | sed -e 's/[^a-zA-Z0-9]/-/g' | tr '[:upper:]' '[:lower:]')
    COMMIT_HASH=$(git rev-parse --short HEAD)
    FEATURE_SUFFIX="-$REF_SLUG-$COMMIT_HASH"
else
    FEATURE_SUFFIX=""
fi

FEATURE_VERSION="$NEXT_VERSION$FEATURE_SUFFIX"

# --- Output ---
# Default behavior: print all versions
if [ "$#" -eq 0 ]; then
    echo "Current Version = $CURRENT_VERSION"
    echo "Feature Version = $FEATURE_VERSION"
    echo "Next Version    = $NEXT_VERSION"
    exit 0
fi

# Flag-based behavior: print specific raw version
case "$1" in
    --current)
        echo "$CURRENT_VERSION"
        ;;
    --next)
        echo "$NEXT_VERSION"
        ;;
    --feature)
        echo "$FEATURE_VERSION"
        ;;
    *)
        echo "Invalid flag. Use --help for usage information." >&2
        exit 1
        ;;
esac
