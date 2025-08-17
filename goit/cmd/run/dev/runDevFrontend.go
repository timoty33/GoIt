package dev

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"time"
)

func RunDevFrontend(configProject utils.ConfigProject) error {
	if configProject.Run.Dev.HotReloadFrontend.Active {
		fmt.Println("Iniciando FrontEnd em modo dev!")
		now := time.Now()
		if file.FileExists("vite.config.js") || file.FileExists("vite.config.ts") {
			return utils.CmdExecuteLog(now, "[FRONTEND]", "vite", configProject.Run.Dev.HotReloadFrontend.ListenPath)
		} else {
			return utils.CmdExecuteLog(now, "[FRONTEND]", "vite", "--port", "3000", "--open", "--cors", "--strictPort", "--base", configProject.Run.Dev.HotReloadFrontend.ListenPath)
		}
	}

	return nil
}
