package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var gpo = &cobra.Command{
	Use:   "gpo",
	Short: "O gpo é o sistema de gerenciamento de plugins para o GOIT",
	Long:  "O gpo é o sistema que o GOIT usa para gerenciar os plugins, ele pode fazer instalações, atualizações, carregar comandos e funções de plugins.",
}

func Execute(versionGpo string) {
	gpo.Version = versionGpo
	if err := gpo.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
