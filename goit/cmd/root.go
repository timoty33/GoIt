package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goit",
	Short: "GoIt é uma ferramenta leve e rápida para scaffolding de projetos Go",
	Long: `GoIt é uma CLI para iniciar projetos em Go com suporte a criação 
de arquivos, estrutura inicial, e integração com banco de dados.`,
}

// Execute chama o rootCmd
func Execute(versionGoit string) {
	rootCmd.Version = versionGoit
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
