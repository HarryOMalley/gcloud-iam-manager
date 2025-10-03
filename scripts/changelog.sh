#!/bin/bash

# --- Changelog Generation Script ---
# Author: Harry O'Malley
#
# This script uses git-cliff to generate a changelog based on conventional commits.

# --- Help Function ---
show_help() {
    echo "Usage: ./changelog.sh [OPTIONS]"
    echo "Generates a changelog using git-cliff."
    echo
    echo "Options:"
    echo "  --from <TAG>      Specify the starting tag for the changelog range."
    echo "  --to <TAG>        Specify the ending tag. If used alone, generates unreleased changes for that tag."
    echo "  --output <FILE>   Write the changelog to the specified file instead of stdout."
    echo "  --help            Show this help message."
}

# --- Argument Parsing ---
FROM_TAG=""
TO_TAG=""
OUTPUT_FILE=""

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --help)
            show_help
            exit 0
            ;;
        --from)
            FROM_TAG="$2"
            shift # past argument
            shift # past value
            ;;
        --to)
            TO_TAG="$2"
            shift # past argument
            shift # past value
            ;;
        --output)
            OUTPUT_FILE="$2"
            shift # past argument
            shift # past value
            ;;
        *)
            echo "Unknown option: $1. Use --help for usage information." >&2
            exit 1
            ;;
    esac
done

# --- Prerequisite Check ---
if ! command -v git-cliff &> /dev/null; then
    echo "Error: git-cliff is not installed. Please install it to continue." >&2
    exit 1
fi

# --- Command Construction ---
CLIFF_CMD="git-cliff"

if [ -n "$FROM_TAG" ] && [ -n "$TO_TAG" ]; then
    CLIFF_CMD="$CLIFF_CMD $FROM_TAG..$TO_TAG"
elif [ -n "$TO_TAG" ]; then
    # If only --to is specified, assume it's for the latest unreleased changes for that tag
    CLIFF_CMD="$CLIFF_CMD --unreleased --tag $TO_TAG"
fi

# --- Execution ---
if [ -n "$OUTPUT_FILE" ]; then
    eval "$CLIFF_CMD" > "$OUTPUT_FILE"
    echo "Changelog written to $OUTPUT_FILE"
else
    eval "$CLIFF_CMD"
fi
