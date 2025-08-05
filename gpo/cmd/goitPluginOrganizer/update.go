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
	var urls []string

	urls, err = utils.LoadJsonListString(filepath.Join(realPathDir, "gpo", "cache.json"))
	if err != nil {
		return fmt.Errorf("erro ao carregar a lista de URLs: %w", err)
	}

	urlGithubPlugin := utils.SearchPlugin(urls, nomePlugin)

	err = UninstallPlugin(nomePlugin)
	if err != nil {
		return fmt.Errorf("erro ao desinstalar o plugin: %s: %w", nomePlugin, err)
	}

	args = []string{urlGithubPlugin}
	err = Install(args)
	if err != nil {
		return fmt.Errorf("erro ao instalar o plugin: %s: %w", nomePlugin, err)
	}

	return nil
}
