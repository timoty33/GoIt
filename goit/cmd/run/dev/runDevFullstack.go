package dev

import (
	"fmt"
	"goit/utils"
)

func RunDevFullstack(configProject utils.ConfigProject) error {
	fmt.Println("Iniciando aplicação em modo de dev!")

	if configProject.Run.Dev.HotReloadBackend.Active {
		go RunDevBackend(configProject)
	}
	if configProject.Run.Dev.HotReloadFrontend.Active {
		go RunDevFrontend(configProject)
	}

	return nil
}
