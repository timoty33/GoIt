package utils

import (
	"fmt"
	"os/exec"
)

func CmdExecute(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao executar comando '%s %v': %w", command, args, err)
	}

	return nil
}
