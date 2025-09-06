package goitPluginOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func UninstallPlugin(nomePlugin string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("erro ao obter o caminho do executável: %v", err)
	}

	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("erro ao resolver o caminho do executável: %v", err)
	}
	realPathDir := filepath.Dir(realPath)

	fullpath := filepath.Join(realPathDir, "plugins", nomePlugin)
	err = os.RemoveAll(fullpath)
	if err != nil {
		return fmt.Errorf("erro ao remover o plugin: %v", err)
	}
	fmt.Println("Plugin removido com sucesso!")

	return nil
}
