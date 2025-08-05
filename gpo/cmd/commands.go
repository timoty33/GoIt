package cmd

func init() {
	gpo.AddCommand(installCmd)
	gpo.AddCommand(uninstallCmd)
	gpo.AddCommand(runCmd)
}
