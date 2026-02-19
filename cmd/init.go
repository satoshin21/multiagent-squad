package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/satoshin21/multiagent/internal/config"
	"github.com/satoshin21/multiagent/internal/embed"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config directory with default layout",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	configDir := config.DefaultConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	files := []struct {
		name    string
		content []byte
	}{
		{"teams.kdl", embed.TeamsKDL},
		{"instruction_leader.md", embed.InstructionLeader},
		{"instruction_member.md", embed.InstructionMember},
	}

	for _, f := range files {
		path := filepath.Join(configDir, f.name)
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("%s already exists, skipping\n", path)
			continue
		}
		if err := os.WriteFile(path, f.content, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", f.name, err)
		}
		fmt.Printf("Created %s\n", path)
	}

	return nil
}
