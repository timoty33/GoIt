package goitPluginOrganizer

import (
	"fmt"
	"goit/utils"
	"os"
	"path/filepath"
)

func Install(args []string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("erro ao obter o caminho do executável: %v", err)
	}

	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("erro ao resolver o caminho do executável: %v", err)
	}
	realPathDir := filepath.Dir(realPath)

	urlGithub := args[0]
	namePlugin := utils.GetPluginNameFromUrl(urlGithub)
	pluginPath := filepath.Join(realPathDir, "plugins", namePlugin)

	err = utils.CmdExecute("git", "clone", urlGithub, pluginPath)
	if err != nil {
		return fmt.Errorf("erro ao clonar o plugin: %v", err)
	}

	return nil
}
