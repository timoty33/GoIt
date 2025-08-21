package cmd

import (
	"goit/cli/gpo/cmd/goitPluginOrganizer"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install <url-do-plugin>",
	Args:  cobra.ExactArgs(1),
	Short: "Instala um plugin do GoIt",

	RunE: func(cmd *cobra.Command, args []string) error {
		return goitPluginOrganizer.Install(args)
	},
}
