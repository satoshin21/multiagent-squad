package runner

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func LaunchGemini(instruction string) error {
	binPath, err := exec.LookPath("gemini")
	if err != nil {
		return fmt.Errorf("gemini not found: %w", err)
	}

	args := []string{"gemini", "--approval-mode", "auto_edit"}
	if instruction != "" {
		args = append(args, instruction)
	}

	return syscall.Exec(binPath, args, os.Environ())
}
