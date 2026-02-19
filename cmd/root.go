package cmd

import (
	"fmt"

	"github.com/satoshin21/multiagent-squad/internal/config"
	"github.com/satoshin21/multiagent-squad/internal/runner"
	"github.com/spf13/cobra"
)

var (
	layoutPath string
	force      bool
)

var rootCmd = &cobra.Command{
	Use:   "multiagent",
	Short: "Multi-agent team launcher for zellij + claude",
	RunE:  runLaunch,
}

func init() {
	rootCmd.Flags().StringVar(&layoutPath, "layout", "", "zellij layout kdl file path (default: ~/.config/teams/teams.kdl)")
	rootCmd.Flags().BoolVar(&force, "force", false, "delete existing session with the same name and create a new one")
	rootCmd.AddCommand(paneCmd)
}

func runLaunch(cmd *cobra.Command, args []string) error {
	layout := layoutPath
	if layout == "" {
		layout = config.DefaultLayoutPath()
	}

	sessionName, err := runner.SessionName(config.DefaultSessionSuffix)
	if err != nil {
		return err
	}

	exists, err := runner.SessionExists(sessionName)
	if err != nil {
		return err
	}
	if exists && !force {
		return fmt.Errorf("session %q already exists; use --force to replace it", sessionName)
	}

	if force && exists {
		fmt.Println("Resetting existing session...")
		runner.DeleteSession(sessionName)
	}

	return runner.CreateSession(sessionName, layout)
}

func Execute() error {
	return rootCmd.Execute()
}
