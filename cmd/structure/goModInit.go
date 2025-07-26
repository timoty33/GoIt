package structure

import (
	"os"
	"os/exec"
)

func GoModInit(nomeProjeto string) error {
	cmd := exec.Command("go", "mod", "init", nomeProjeto)
	cmd.Dir = nomeProjeto
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
