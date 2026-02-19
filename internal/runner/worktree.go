package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func EnsureWorktree(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path %q: %w", path, err)
	}
	if _, err := os.Stat(absPath); err == nil {
		return absPath, nil
	}
	if err := exec.Command("git", "worktree", "add", absPath).Run(); err != nil {
		return "", fmt.Errorf("failed to add worktree at %q: %w", absPath, err)
	}
	return absPath, nil
}
