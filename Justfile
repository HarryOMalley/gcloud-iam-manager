# Justfile for the gcloud-iam-manager project

# Build and run the application
run:
    go run ./cmd/gcloud-iam-manager

# Build the application binary
build:
    go build ./cmd/gcloud-iam-manager

# Run all tests
test:
    go test ./...

# Lint all files using the configured pre-commit hooks
lint:
    pre-commit run --all-files

# Display version information using our script
version:
    ./scripts/version.sh

# Tail the latest CI run for the current branch
ci-tail:
    ./scripts/ci-tail.sh
