package cmd

import (
	"fmt"
	"goit/gpo/cmd/goitPluginOrganizer"
	"goit/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <nome-do-plugin> [args...]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Executa um plugin do GoIt",
	RunE: func(cmd *cobra.Command, args []string) error {
		binPath, argsCommand, err := goitPluginOrganizer.CommandRun(args)
		if err != nil {
			return fmt.Errorf("erro ao preparar o comando: %w", err)
		}

		configProject, configPaths, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("erro ao carregar configuração: %w", err)
		}

		configComplete := utils.Config{
			Project: configProject,
			Paths:   configPaths,
		}

		err = utils.CmdExecuteWithJSONInput(binPath, configComplete, argsCommand...) // o '...' desempacota o slice argsCommand
		if err != nil {
			return fmt.Errorf("erro ao executar o comando: %w", err)
		}

		return nil
	},
}
