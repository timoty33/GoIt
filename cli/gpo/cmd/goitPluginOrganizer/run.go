package goitPluginOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func CommandRun(args []string) (string, []string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", nil, fmt.Errorf("erro ao obter o caminho do executável: %w", err)

	}
	exeDir := filepath.Dir(exePath)

	goos := runtime.GOOS
	goarch := runtime.GOARCH
	path := fmt.Sprintf("%s_%s", goos, goarch)

	pluginName := args[0]
	argsCommand := args[1:]

	pluginBinPath := filepath.Join(exeDir, "plugins", pluginName, "bin", path)
	switch goos {
	case "windows":
		pluginBinPath = filepath.Join(pluginBinPath, pluginName+".exe")
	default:
		pluginBinPath = filepath.Join(pluginBinPath, pluginName)
	}

	if _, err := os.Stat(pluginBinPath); os.IsNotExist(err) {
		return "", nil, fmt.Errorf("binário do plugin não encontrado: %s", pluginBinPath)

	}

	return pluginBinPath, argsCommand, nil
}
