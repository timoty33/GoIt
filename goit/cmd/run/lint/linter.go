package lint

import (
	"goit/utils"
	"os/exec"
)

// ================ Biome Setup ================
func IsBiomeAvaible() bool {
	cmd := exec.Command("npx", "biome", "--version")
	err := cmd.Run()
	return err == nil
}

func InstallBiomeNpm() error {
	return utils.CmdExecute("npm", "install", "--save-dev", "biome")
}

func BiomeInit() error {
	return utils.CmdExecute("npx", "biome", "init")
}

// ======== Biome Linter and Formatter =========
func BiomeLint(lintFrontend string) error {
	return utils.CmdExecute("npx", "biome", "lint", lintFrontend)
}

func BiomeLintApply(lintFrontend string) error {
	return utils.CmdExecute("npx", "biome", "lint", lintFrontend, "--apply")
}

func BiomeFormatter(lintFrontend string) error {
	return utils.CmdExecute("npx", "biome", "format", lintFrontend)
}

// =========== Static Check Setup ==============
func IsStaticCheckAvaible() bool {
	cmd := exec.Command("staticcheck", "--version")
	err := cmd.Run()
	return err == nil
}

func InstallStaticCheck() error {
	return utils.CmdExecute("go", "install", "honnef.co/go/tools/cmd/staticcheck@latest")
}

// ======= Static Linter and Formatter =========
func StaticLint(path string) error {
	return utils.CmdExecute("staticcheck", path)
}

func GoFmt(path string) error {
	return utils.CmdExecute("go", "fmt", path)
}
