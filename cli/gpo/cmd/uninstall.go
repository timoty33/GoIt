package cmd

import (
	"goit/cli/gpo/cmd/goitPluginOrganizer"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Desinstala um plugin do GoIt",
	Long:  `Desinstala um plugin do GoIt, removendo-o da pasta plugins.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nomePlugin := args[0]
		return goitPluginOrganizer.UninstallPlugin(nomePlugin)
	},
}
