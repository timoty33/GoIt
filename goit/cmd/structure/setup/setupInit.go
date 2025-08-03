package setup

import (
	"goit/utils"
)

func GoModInit(projectPath, nomeProjeto string) error {
	return utils.CmdExecuteInDir(projectPath, "go", "mod", "init", nomeProjeto)
}

func PythonInit(projectPath string) error {
	// Cria um ambiente virtual chamado 'venv' dentro do diretório do projeto.
	return utils.CmdExecuteInDir(projectPath, "python", "-m", "venv", "venv")
}

func NodeInit(projectPath string) error {
	return utils.CmdExecuteInDir(projectPath, "npm", "init", "-y")
}

func TsInit(projectPath string) error {
	return utils.CmdExecuteInDir(projectPath, "npx", "tsc", "--init")
}
