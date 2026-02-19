package cmd

import (
	"fmt"
	"os"

	"github.com/satoshin21/multiagent/internal/runner"
	"github.com/spf13/cobra"
)

var (
	instruction             string
	claudeWithoutInstruction bool
)

var paneCmd = &cobra.Command{
	Use:   "pane <worktree-path> <branch>",
	Short: "Move to worktree and optionally launch claude with instruction",
	Args:  cobra.ExactArgs(2),
	RunE:  runPane,
}

func init() {
	paneCmd.Flags().StringVar(&instruction, "instruction", "", "instruction markdown file path")
	paneCmd.Flags().BoolVar(&claudeWithoutInstruction, "claude-without-instruction", false, "launch claude without instruction")
}

func runPane(cmd *cobra.Command, args []string) error {
	worktreePath := args[0]
	branch := args[1]

	wtPath, err := runner.EnsureWorktree(worktreePath, branch)
	if err != nil {
		return fmt.Errorf("failed to ensure worktree: %w", err)
	}

	if err := os.Chdir(wtPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", wtPath, err)
	}

	if claudeWithoutInstruction {
		return runner.LaunchClaude("")
	}

	if instruction == "" {
		return nil
	}

	return runner.LaunchClaude(instruction)
}
