package cmd

import (
	"fmt"
	"goit/goit/cmd/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [--project/--path] [chave] [valor]",
	Short: "Configurações do projeto",
	Long:  "Você pode modificar campos específicos das configurações do projeto ou dos caminhos/paths",
	Args:  cobra.ExactArgs(2), // chave e valor
	RunE: func(cmd *cobra.Command, args []string) error {
		var key = args[0]
		var value = args[1]

		err := config.UpdateConfigKey(configFile, key, value)
		if err != nil {
			return fmt.Errorf("erro ao modificar configs: %w", err)
		}

		return nil

	},
}
