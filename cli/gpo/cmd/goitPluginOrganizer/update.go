package goitPluginOrganizer

import (
	"fmt"
	"goit/utils"
	"os"
	"path/filepath"
)

func UpdatePlugin(args []string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("erro ao obter o caminho do executável: %v", err)
	}

	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("erro ao resolver o caminho do executável: %v", err)
	}
	realPathDir := filepath.Dir(realPath)

	nomePlugin := args[0]
	err = utils.CmdExecuteInDir(filepath.Join(realPathDir, "plugins", nomePlugin), "git", "pull")
	if err != nil {
		return fmt.Errorf("erro ao atualizar o plugin: %w", err)
	}
	fmt.Println("Plugin atualizado com sucesso!")

	return nil
}
