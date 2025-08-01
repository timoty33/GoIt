package setup

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

func PythonInit(nomeProjeto string) error {
	cmd := exec.Command("python", "-m", "venv", "venv")
	cmd.Dir = nomeProjeto
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func NodeInit(nomeProjeto string) error {
	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = nomeProjeto
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TsInit(nomeProjeto string) error {
	// Executa "npx tsc --init"
	cmd := exec.Command("npx", "tsc", "--init")
	cmd.Dir = nomeProjeto
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}
