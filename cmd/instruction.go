package cmd

import (
	"fmt"
	"os"

	"github.com/satoshin21/multiagent/internal/runner"
	"github.com/spf13/cobra"
)

var (
	instruction string
)

var paneCmd = &cobra.Command{
	Use:   "pane <worktree>",
	Short: "Move to worktree and optionally launch claude with instruction",
	Args:  cobra.ExactArgs(1),
	RunE:  runPane,
}

func init() {
	paneCmd.Flags().StringVar(&instruction, "instruction", "", "instruction markdown file path")
}

func runPane(cmd *cobra.Command, args []string) error {
	worktreePath := args[0]

	wtPath, err := runner.EnsureWorktree(worktreePath)
	if err != nil {
		return fmt.Errorf("failed to ensure worktree: %w", err)
	}

	if err := os.Chdir(wtPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", wtPath, err)
	}

	if instruction == "" {
		return nil
	}

	return runner.LaunchClaude(instruction)
}
