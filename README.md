# multiagent

Multi-agent team launcher for [Zellij](https://zellij.dev/) + Claude Code.
Creates a Zellij session with a team layout and launches Claude in each git worktree.

## Prerequisites

- **Go** 1.21+
- **Zellij** — https://zellij.dev/documentation/installation.html
- **Claude Code** — https://docs.anthropic.com/en/docs/claude-code
## Install

```bash
go install github.com/satoshin21/multiagent@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Setup

```bash
multiagent init
```

Writes the default Zellij layout file to `~/.config/teams/teams.kdl`.
Skipped if the file already exists.

To customize the layout, edit `~/.config/teams/teams.kdl` directly.

## Usage

### launch

Create a Zellij session with the team layout. Any existing session with the same name is automatically reset.

```bash
multiagent launch
multiagent launch --layout /path/to/custom.kdl
```

Running `multiagent` without arguments is equivalent to `multiagent launch`.

### instruction

Change into the specified worktree and launch Claude. Intended to be called from Zellij pane commands.

```bash
multiagent instruction --worktree team-member/sei --instruction ~/.config/teams/instruction_sei.md
multiagent instruction --worktree team-member/shin
```

| Flag | Description |
|------|-------------|
| `--worktree` | Worktree directory path (required) |
| `--instruction` | Path to instruction markdown file (optional) |

### init

Write the default layout file to `~/.config/teams/teams.kdl`.

```bash
multiagent init
```

## License

See repository license.
