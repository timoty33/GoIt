package structure

import (
	"fmt"
	"goit/cmd"
	"goit/cmd/file"
	"os"
	"os/exec"
)

const (
	permPasta   = 0755
	permArquivo = 0644
)

func GetModGin(projectPath string) error {
	cmd := exec.Command("go", "get", "github.com/gin-gonic/gin")
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("falha ao executar 'go get': %w", err)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func StructureGin(nomeProjeto string) error {
	err := os.Mkdir(nomeProjeto, permPasta)
	if err != nil {
		return fmt.Errorf("erro ao criar pasta do projeto %s: %w", nomeProjeto, err)
	}

	// Renderiza os templates
	templates, err := file.PercorrerDiretorio("../templates")
	if err != nil {
		return fmt.Errorf("erro ao percorrer templates: %w", err)
	}

	err = cmd.RenderTemplates(templates, cmd.TemplateData{ProjectName: nomeProjeto}, nomeProjeto)
	if err != nil {
		return fmt.Errorf("erro ao renderizar templates: %w", err)
	}

	return nil
}
