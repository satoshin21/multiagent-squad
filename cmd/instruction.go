package cmd

import (
	"fmt"
	"os"

	"github.com/satoshin21/multiagent/internal/runner"
	"github.com/spf13/cobra"
)

var (
	instructionFile string
	worktreePath    string
)

var instructionCmd = &cobra.Command{
	Use:   "instruction",
	Short: "Move to worktree and launch claude with instruction",
	RunE:  runInstruction,
}

func init() {
	instructionCmd.Flags().StringVar(&instructionFile, "instruction", "", "instruction markdown file path")
	instructionCmd.Flags().StringVar(&worktreePath, "worktree", "", "worktree directory path (required)")
	_ = instructionCmd.MarkFlagRequired("worktree")
}

func runInstruction(cmd *cobra.Command, args []string) error {
	wtPath, err := runner.EnsureWorktree(worktreePath)
	if err != nil {
		return fmt.Errorf("failed to ensure worktree: %w", err)
	}

	if err := os.Chdir(wtPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", wtPath, err)
	}

	return runner.LaunchClaude(instructionFile)
}
