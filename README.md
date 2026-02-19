# multiagent-squad

Multi-agent team launcher for [Zellij](https://zellij.dev/) + [Claude Code](https://docs.anthropic.com/en/docs/claude-code).

Creates a Zellij session with a multi-pane team layout, launching isolated Claude Code instances in separate git worktrees. Chief (commander) delegates tasks to Member-Braze / Member-Storm / Member-Frost (team members), enabling parallel development through a team-of-agents pattern.

## Prerequisites

- **Go** 1.21+
- **Zellij** — https://zellij.dev/documentation/installation.html
- **Claude Code** — https://docs.anthropic.com/en/docs/claude-code

## Install

```bash
go install github.com/satoshin21/multiagent-squad@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Setup

```bash
multiagent init
```

Creates the following default config files in `~/.config/multiagent-squad/` (existing files are skipped).

| File | Description |
|------|-------------|
| `squad.kdl` | Zellij layout definition |
| `instruction_leader.md` | Claude Code instruction for Chief (commander) |
| `instruction_member.md` | Claude Code instruction for members (Member-Braze / Member-Storm / Member-Frost) |

Edit these files directly to customize the layout or instructions.

## Usage

### Default (launch)

```bash
multiagent
```

Running without arguments creates a Zellij session with the default layout. The session is named `<current-directory-name>-teams`.

| Flag | Description |
|------|-------------|
| `--layout <path>` | Path to a custom layout file (default: `~/.config/multiagent-squad/squad.kdl`) |
| `--force` | Delete and recreate the session if one with the same name already exists |

```bash
multiagent --layout /path/to/custom.kdl
multiagent --force
```

### pane

Used internally as a Zellij pane command. Ensures the git worktree exists and launches Claude Code.

```bash
multiagent pane <worktree-path> <branch> [flags]
```

| Flag | Description |
|------|-------------|
| `--instruction <text>` | Instruction text passed to Claude Code |
| `--claude-without-instruction` | Launch Claude Code without any instruction |

### init

Creates default config files in `~/.config/multiagent-squad/`.

```bash
multiagent init
```

## Architecture

### Default Layout

The default layout consists of two tabs: `agent-teams` and `review`.

```
agent-teams tab
┌──────────┬───────────────────────┐
│          │     member-braze      │
│          ├───────────────────────┤
│  chief   │     member-storm      │
│   (30%)  ├───────────────────────┤
│          │     member-frost      │
└──────────┴───────────────────────┘

review tab
┌──────────────┬──────────────────┐
│   terminal   │ claude (no inst) │
└──────────────┴──────────────────┘
```

### git worktree

Each agent operates in an independent git worktree at `../multiagent_worktrees/team-member/<name>`, allowing parallel work without branch conflicts.

### Inter-agent Communication

Agents communicate by sending messages to other panes via Zellij's `send-to-pane`.

| Pane ID | Agent |
|---------|-------|
| 0 | chief |
| 1 | member-braze |
| 2 | member-storm |
| 3 | member-frost |

## License

See repository license.
