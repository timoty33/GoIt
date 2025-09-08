package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/process"
)

func ErrorfatUptime(start time.Time) string {
	uptime := time.Since(start) // retorna time.Duration

	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// CmdExecute executa um comando no diretório de trabalho atual.
// (Esta é uma suposição de como sua função existente pode ser)
func CmdExecute(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CmdExecuteLog(timeInit time.Time, logPrefix, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	pid := cmd.Process.Pid
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return err
	}

	getProcessStatus := func() any {
		memInfo, _ := p.MemoryInfo()
		memMB := memInfo.RSS / 1024 / 1024

		return memMB
	}

	printWithPrefix := func(r io.Reader, prefix string) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Printf("\033[1;32m%s %s | RAM: %v mb | Rodando a %s \033[0m \n  |->> %s\n\n", prefix, time.Now().Format("15:04:05.000"), getProcessStatus(), ErrorfatUptime(timeInit), scanner.Text())
		}
	}

	printWithPrefixError := func(r io.Reader, prefix string) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Printf("\033[1;31m[ERRO] %s %s | RAM: %v mb | Rodando a %s \033[0m \n  |->> %s\n\n", prefix, time.Now().Format("15:04:05.000"), getProcessStatus(), ErrorfatUptime(timeInit), scanner.Text())
		}
	}

	go printWithPrefix(stdout, logPrefix)
	go printWithPrefixError(stderr, logPrefix)

	return cmd.Wait()
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
