---
name: NetScanner
description: Asynchronous parsing and execution of lsof to identify LISTEN ports
---

# Skill: NetScanner

## Objective
Acquire in real-time the list of listening (LISTEN) TCP and UDP ports on the macOS system, asynchronously so as not to block the Bubble Tea UI.

## Implementation
- **Target macOS command:** `lsof -iTCP -iUDP -sTCP:LISTEN -P -n +c 0`
  - `-P`: Do not convert port numbers to service names.
  - `-n`: Do not resolve hostnames (speed gain).
  - `+c 0`: Display the full command name without truncation.
- **Parsing:** Read the standard output (stdout) line by line, ignore the header, and extract: PID, COMM, TYPE, NODE, NAME.
- **Translation:** Map common command names to readable descriptions (e.g. `node` -> "JS Dev Server", `postgres` -> "PostgreSQL Database").
- **Bubble Tea:** Return results via a `tea.Msg` containing a list of `system.PortInfo` structures.
