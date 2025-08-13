package cmd

import (
	"fmt"
	"goit/goit/cmd/run/lint"
	"goit/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Inicia um super projeto, podendo ser de diferentes modos",
	Long:  "Inicia um super projeto criado com GoIt, e pode ter diferentes modos, como 'dev', 'test' ou 'linter'",

	RunE: func(cmd *cobra.Command, args []string) error {
		// carregando as configs
		configProject, _, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("erro ao carregar configurações, %w", err)
		}

		commandArg := args[0]
		switch commandArg {
		case "linter":

			// passa o linter no front
			if configProject.ProjectType == "FullStack" || configProject.ProjectType == "Frontend" {
				err := lint.RunBiome(configProject)
				if err != nil {
					return fmt.Errorf("erro ao usar o biome: %w", err)
				}
				fmt.Println("Biome feito com sucesso!")
			}

			if configProject.ProjectType == "Backend" || configProject.ProjectType == "FullStack" {
				switch configProject.ProgrammingLanguage {
				case "Go":
					if err := lint.RunStaticFmt(configProject); err != nil {
						return fmt.Errorf("erro ao rodar linter para go: %w", err)
					}
				}
			}
		}

		return nil
	},
}
