package setup

import (
	"goit/utils"
)

func GoModInit(projectPath, nomeProjeto string) error {
	return utils.CmdExecuteInDir(projectPath, "go", "mod", "init", nomeProjeto)
}

func NodeInit(projectPath string) error {
	return utils.CmdExecuteInDir(projectPath, "npm", "init", "-y")
}

func TsInit(projectPath string) error {
	return utils.CmdExecuteInDir(projectPath, "npx", "tsc", "--init")
}
