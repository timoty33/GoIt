package setup

import (
	"goit/utils"
)

func GoModInit(nomeProjeto string) error {
	return utils.CmdExecute("go", "mod", "init", nomeProjeto)
}

func PythonInit(nomeProjeto string) error {
	return utils.CmdExecute("python", "-m", "venv", nomeProjeto)
}

func NodeInit(nomeProjeto string) error {
	return utils.CmdExecute("npm", "init", "-y")
}

func TsInit(nomeProjeto string) error {
	return utils.CmdExecute("npx", "tsc", "--init")
}
