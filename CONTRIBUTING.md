# Contributing to gcloud-iam-manager

We welcome contributions to this project. Please follow these guidelines to ensure a smooth and effective development process.

## Development Methodology

This project follows a strict **Test-Driven Development (TDD)** workflow. All new features or bug fixes must be accompanied by tests.

1.  **Write a Failing Test**: Before writing any implementation code, add a test that specifies the desired functionality and fails because the functionality does not exist.
2.  **Write Code to Pass**: Write the simplest possible code to make the test pass.
3.  **Refactor**: Improve the code's structure and clarity while ensuring all tests continue to pass.

For more detailed instructions, please refer to the `AI_TDD.md` guide.

## Commit Conventions

All commit messages must adhere to the **Conventional Commits** specification. This is essential for automating changelog generation and versioning.

*   **Format**: `<type>(<scope>): <subject>`
*   **Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`.

For detailed instructions on crafting commit messages, refer to the `AI_COMMITS.md` guide.

## Branching Strategy

All new work should be done on a feature branch. Branch names should be descriptive and prefixed with the type of work being done.

*   **Features**: `feat/add-role-picker`
*   **Bug Fixes**: `fix/user-list-panic`
*   **Documentation**: `docs/update-readme`

## Merge Requests

When you are ready to merge your work, open a Merge Request (or Pull Request). The description should be clear and follow the template outlined in `AI_MR_DESCRIPTION.md`.

## Versioning and Tagging

We use `git-cliff` to generate a changelog and determine the next semantic version number based on the commit history. Tags are created without a `v` prefix (e.g., `0.1.0`, `1.0.0`).
