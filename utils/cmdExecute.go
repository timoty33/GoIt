package utils

import (
	"encoding/json"
	"fmt"
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

func CmdExecuteWithJSONInput(name string, jsonInput any, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Serializa a struct
	go func() {
		defer stdin.Close()
		if err = json.NewEncoder(stdin).Encode(jsonInput); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao codificar JSON: %v\n", err)
		}
	}()

	return cmd.Run()
}
