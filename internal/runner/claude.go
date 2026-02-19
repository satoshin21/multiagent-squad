package runner

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/satoshin21/multiagent/internal/config"
)

func LaunchClaude(instruction string) error {
	binPath, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude not found: %w", err)
	}

	args := []string{"claude", "--permission-mode", "acceptEdits"}

	if instruction != "" {
		expanded := config.ExpandTilde(instruction)
		if _, err := os.Stat(expanded); err == nil {
			args = append(args, instruction)
		}
	}

	return syscall.Exec(binPath, args, os.Environ())
}
