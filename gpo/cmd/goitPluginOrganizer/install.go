package goitPluginOrganizer

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var InstallGpo = &cobra.Command{
	Use:   "install",
	Args:  cobra.ExactArgs(1),
	Short: "Instala um plugin do GoIt",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Obtém o caminho do executável atual
		// Isso é necessário para garantir que o plugin seja instalado no diretório correto
		exePath, err := os.Executable()
		if err != nil {
			panic(err)
		}

		// Resolve symlinks e retorna o caminho absoluto
		realPath, err := filepath.EvalSymlinks(exePath)
		if err != nil {
			panic(err)
		}

		realPathDir := filepath.Dir(realPath)

		urlGithub := args[0]
		namePlugin := utils.GetPluginNameFromUrl(urlGithub)

		utils.CmdExecute("git", "clone", urlGithub, filepath.Join(realPathDir, "plugins", namePlugin))

		_, err = file.ReadFile(filepath.Join(realPathDir, "plugins", namePlugin, "plugin.go"))
		if err != nil {
			return fmt.Errorf("erro ao ler o arquivo plugin.go: %v", err)
		}

		return nil
	},
}
