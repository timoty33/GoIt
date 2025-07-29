package structure

import (
	"fmt"
	"goit/cmd/structure/setup"
)

func Setup(nomeProjeto, linguagemProjeto string) error {
	fmt.Println("Iniciando o setup do projeto:", nomeProjeto)

	switch linguagemProjeto {
	case "Go":
		if err := setup.GoModInit(nomeProjeto); err != nil {
			return fmt.Errorf("erro ao inicializar o módulo Go: %w", err)
		}
	case "Python":
		if err := setup.PythonInit(nomeProjeto); err != nil {
			return fmt.Errorf("erro ao inicializar o projeto Python (venv): %w", err)
		}
	case "JavaScript", "TypeScript":
		if err := setup.NodeInit(nomeProjeto); err != nil {
			return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
		}
	}

	return nil
}
