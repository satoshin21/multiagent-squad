package runner

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func SessionExists(sessionName string) (bool, error) {
	cmd := exec.Command("zellij", "list-sessions", "--short")
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("list-sessions: %w", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == sessionName {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func SessionName(suffix string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}
	return filepath.Base(cwd) + "-" + suffix, nil
}

func DeleteSession(sessionName string) {
	cmd := exec.Command("zellij", "delete-session", "--force", sessionName)
	_ = cmd.Run()
}

func CreateSession(sessionName string, layout string) error {
	binPath, err := exec.LookPath("zellij")
	if err != nil {
		return fmt.Errorf("zellij not found: %w", err)
	}

	args := []string{"zellij", "--session", sessionName, "--new-session-with-layout", layout}
	return syscall.Exec(binPath, args, os.Environ())
}
