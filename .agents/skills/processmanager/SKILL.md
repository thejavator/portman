---
name: ProcessManager
description: Secure execution of termination signals (kill) and favorite filtering
---

# Skill: ProcessManager

## Objective
Allow the termination of one or more processes selected in the TUI, respecting security rules (favorites).

## Implementation
- **Target commands:** `kill -15 <pid>` (Graceful) and `kill -9 <pid>` (Force).
- **Validation:** Before each kill, check if the targeted process is part of the "Favorites" list (protected processes loaded from `config.go`).
- **Errors:** If a `kill` fails due to permissions (root process), return an explicit error via `tea.Msg` to display it as an alert in the TUI.
- **Sudo:** Rely on the privilege elevation of the parent process (portman) to ensure that `kill` has the necessary rights.
