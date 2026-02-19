package cmd

import (
	"fmt"

	"github.com/satoshin21/multiagent/internal/config"
	"github.com/satoshin21/multiagent/internal/runner"
	"github.com/spf13/cobra"
)

var (
	layoutPath string
	force      bool
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Create a zellij session with team layout",
	RunE:  runLaunch,
}

func init() {
	launchCmd.Flags().StringVar(&layoutPath, "layout", "", "zellij layout kdl file path (default: ~/.config/teams/teams.kdl)")
	launchCmd.Flags().BoolVar(&force, "force", false, "delete existing session with the same name and create a new one")
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
