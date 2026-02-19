package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "multiagent",
	Short: "Multi-agent team launcher for zellij + claude",
	RunE: func(cmd *cobra.Command, args []string) error {
		return launchCmd.RunE(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
	rootCmd.AddCommand(instructionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
