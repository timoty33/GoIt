package dev

import (
	"fmt"
	"goit/goit/cmd/run/dev/watchexec"
	"goit/utils"
	"time"
)

func RunDevBackend(configProject utils.ConfigProject) error {
	fmt.Println("Iniciando BackEnd em modo dev!")
	now := time.Now()

	// se o watchexec estiver instalado, daí criamos
	// o comando e executamos ele
	if watchexec.WatchexecInstalled() {
		// gera o comando
		commandWatch := []string{"watchexec", "-r"}
		for _, i := range configProject.Run.Dev.Ignore {
			commandWatch = append(commandWatch, "--ignore")
			commandWatch = append(commandWatch, i)
		}
		commandWatch = append(commandWatch, configProject.Run.Dev.InitCommandBackend)

		// executa ele
		fmt.Println(commandWatch)
		utils.CmdExecuteLog(now, "[BACKEND]", commandWatch[0], commandWatch[1:]...)
	}

	return nil
}
