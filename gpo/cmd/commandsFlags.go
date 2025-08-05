package cmd

import (
	"goit/gpo/cmd/goitPluginOrganizer"
)

func init() {
	gpo.AddCommand(goitPluginOrganizer.InstallGpo)
	gpo.AddCommand(runCmd)
}
