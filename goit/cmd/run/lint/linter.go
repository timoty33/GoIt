package lint

import (
	"goit/utils"
	"os/exec"
)

// ================ Biome Setup ================
func isBiomeAvaible() bool {
	cmd := exec.Command("npx", "biome", "--version")
	err := cmd.Run()
	return err == nil
}

func installBiomeNpm() error {
	return utils.CmdExecute("npm", "i", "-D", "-E", "@biomejs/biome")
}

func biomeInit() error {
	return utils.CmdExecute("npx", "@biomejs/biome", "init")
}

// ======== Biome Linter and Formatter =========
func biomeRunApply(configProject utils.ConfigProject) error {
	return utils.CmdExecute("npx", "@biomejs/biome", "check", "--write", configProject.Run.Lint.LintFrontEnd)
}

func biomeRun(configProject utils.ConfigProject) error {
	return utils.CmdExecute("npx", "@biomejs/biome", "check", configProject.Run.Lint.LintFrontEnd)
}

// =========== Static Check Setup ==============
func isStaticCheckAvaible() bool {
	cmd := exec.Command("staticcheck", "--version")
	err := cmd.Run()
	return err == nil
}

func installStaticCheck() error {
	return utils.CmdExecute("go", "install", "honnef.co/go/tools/cmd/staticcheck@latest")
}

// ======= Static Linter and Formatter =========
func staticLint(path string) error {
	return utils.CmdExecute("staticcheck", path)
}

func goFmt(path string) error {
	return utils.CmdExecute("go", "fmt", path)
}
