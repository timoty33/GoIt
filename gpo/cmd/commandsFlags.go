package cmd

import (
	"goit/gpo/cmd/goitPluginOrganizer"
	// goit:add-imports-here
)

func init() {
	gpo.AddCommand(goitPluginOrganizer.InstallGpo)

	// start functions from plugins
	// goit:add-plugins-here
}
