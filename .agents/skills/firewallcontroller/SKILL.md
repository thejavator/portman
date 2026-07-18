---
name: FirewallController
description: Blocking/Unblocking network traffic via pfctl
---

# Skill: FirewallController

## Objective
Manage inbound and outbound network traffic for a given port by directly interacting with the macOS Packet Filter (`pfctl`).

## Implementation
- **Mechanism:** Use `pfctl` anchors or generate a temporary rule file.
  - Port block example: `echo "block drop in proto tcp from any to any port <PORT>" | sudo pfctl -a com.portman.block -f -`
- **Unblocking:** Flush the rules of the specific anchor or port.
  - Example: `sudo pfctl -a com.portman.block -F rules`
- **Reloading:** Silently enable `pf` if necessary (`sudo pfctl -e` and mask the error if it is already active).
- **Bubble Tea:** Execute these commands as asynchronous `tea.Cmd` and return the status (success/failure) to the UI.
