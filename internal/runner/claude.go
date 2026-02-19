package runner

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/satoshin21/multiagent/internal/config"
)

func LaunchClaude(instructionFile string) error {
	binPath, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude not found: %w", err)
	}

	args := []string{"claude", "--permission-mode", "acceptEdits"}

	if instructionFile != "" {
		expanded := config.ExpandTilde(instructionFile)
		if _, err := os.Stat(expanded); err == nil {
			args = append(args, "Read "+instructionFile+" and understand your role.")
		}
	}

	return syscall.Exec(binPath, args, os.Environ())
}
