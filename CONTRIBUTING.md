# Contributing to Portman

First off, thank you for considering contributing to Portman! We aim to build the best terminal-based port and process manager for macOS.

## Code of Conduct

By participating in this project, you agree to abide by our standard code of conduct. Be respectful, inclusive, and professional.

## How to Contribute

### 1. Reporting Bugs

If you find a bug, please open an issue and include:
- Your macOS version.
- The version of Go you are using.
- Steps to reproduce the bug.
- Any relevant logs or output.

### 2. Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. Please provide:
- A clear and descriptive title.
- A detailed description of the proposed feature.
- Why this feature would be useful to most Portman users.

### 3. Pull Requests

We gladly accept Pull Requests! Please follow this process:

1. **Fork the repo** and create your branch from `main`.
2. **Write idiomatic Go**: Follow standard Go formatting conventions.
3. **Format your code**: Run `go fmt ./...` before committing.
4. **Test your code**: Ensure the application builds cleanly with `go build`.
5. **Commit Messages**: Use clear, concise commit messages (e.g., `feat: add IPv6 support in scanner` or `fix: handle missing pfctl permissions`).
6. **Open a PR**: Describe your changes in detail, link to any relevant issues, and submit for review.

## Architecture Overview

- **`tui/`**: Contains the Bubble Tea UI models, update loops, and view rendering.
- **`system/`**: Encapsulates all OS-level commands (`lsof`, `ps`, `pfctl`, `kill`).
- **`config/`**: Handles loading and saving persistent settings to `~/.config/portman/config.json`.

Happy coding! 🚀
