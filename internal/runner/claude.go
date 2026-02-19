package runner

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func LaunchClaude(instruction string) error {
	binPath, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude not found: %w", err)
	}

	args := []string{"claude", "--permission-mode", "acceptEdits"}

	if instruction != "" {
		args = append(args, instruction)
	}

	return syscall.Exec(binPath, args, os.Environ())
}
