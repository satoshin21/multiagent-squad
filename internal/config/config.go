package config

import (
	"os"
	"path/filepath"
	"strings"
)

const DefaultSessionSuffix = "teams"

func ExpandTilde(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	return filepath.Join(home, path[1:])
}

func DefaultConfigDir() string {
	return ExpandTilde("~/.config/teams")
}

func DefaultLayoutPath() string {
	return filepath.Join(DefaultConfigDir(), "teams.kdl")
}
