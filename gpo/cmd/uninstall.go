package cmd

import (
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Desinstala um plugin do GoIt",
	Long:  `Desinstala um plugin do GoIt, removendo-o da pasta plugins.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil

	},
}
