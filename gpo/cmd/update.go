package cmd

import (
	"fmt"
	"goit/gpo/cmd/goitPluginOrganizer"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza um plugin do GoIt",
	Long:  "Deleta e reinstala um plugin do GoIt",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		err := goitPluginOrganizer.UpdatePlugin(args)
		if err != nil {
			return fmt.Errorf("erro ao atualizar o plugin: %w", err)
		}
		return nil
	},
}
