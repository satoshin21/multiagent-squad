package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func EnsureWorktree(path, branch string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path %q: %w", path, err)
	}
	if _, err := os.Stat(absPath); err == nil {
		return absPath, nil
	}
	// 実体が削除済みだが登録が残っている古い worktree を除去
	if err := exec.Command("git", "worktree", "add", "-B", branch, absPath).Run(); err != nil {
		return "", fmt.Errorf("failed to add worktree at %q with branch %q: %w", absPath, branch, err)
	}
	return absPath, nil
}
