package utils

import (
	"os"
	"os/exec"
)

// CmdExecute executa um comando no diretório de trabalho atual.
// (Esta é uma suposição de como sua função existente pode ser)
func CmdExecute(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CmdExecuteInDir executa um comando no diretório especificado.
func CmdExecuteInDir(dir, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir // Define o diretório de trabalho para o comando
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
