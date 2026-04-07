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
	out, err := cmd.CombinedOutput()
	return parseSessionExistsOutput(sessionName, out, err)
}

func parseSessionExistsOutput(sessionName string, out []byte, cmdErr error) (bool, error) {
	output := string(out)
	if cmdErr != nil {
		if strings.Contains(output, "No active zellij sessions found") {
			return false, nil
		}
		return false, fmt.Errorf("list-sessions: %w", cmdErr)
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == sessionName {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
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
