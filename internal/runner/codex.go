package runner

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func LaunchCodex(instruction string) error {
	binPath, err := exec.LookPath("codex")
	if err != nil {
		return fmt.Errorf("codex not found: %w", err)
	}

	args := []string{"codex", "--full-auto"}
	if instruction != "" {
		args = append(args, instruction)
	}

	return syscall.Exec(binPath, args, os.Environ())
}
