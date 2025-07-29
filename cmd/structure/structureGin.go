package structure

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
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

func CreateStructure(nomeProjeto, linguagem, framework, tipoProjeto string) (utils.ConfigPaths, error) {
	var configPaths utils.ConfigPaths

	err := os.Mkdir(nomeProjeto, permPasta)
	if err != nil {
		return configPaths, fmt.Errorf("erro ao criar pasta do projeto %s: %w", nomeProjeto, err)
	}

	if tipoProjeto == "Backend" || tipoProjeto == "FullStack" {

		// Renderiza os templates (backend, o front será criado com Vite)
		templates, err := file.PercorrerDiretorio("../templates/" + linguagem + "/" + framework)
		if err != nil {
			return configPaths, fmt.Errorf("erro ao percorrer templates: %w", err)
		}

		err = RenderTemplates(templates, TemplateData{ProjectName: nomeProjeto}, nomeProjeto)
		if err != nil {
			return configPaths, fmt.Errorf("erro ao renderizar templates: %w", err)
		}
	}

	switch linguagem {
	case "Go":
		switch framework {
		case "Gin":
			configPaths = utils.ConfigPaths{
				RoutesFile: "internal/routes/routes.go",

				HandlersFile:   "internal/handlers/handlers.go",
				HandlersFolder: "internal/handlers",

				MiddlewaresFolder: "internal/middlewares",
				MiddlewaresFile:   "internal/middlewares/middlewares.go",

				DtoFolder: "internal/dto",
				DtoFile:   "internal/dto/dto.go",

				ModelsFolder: "internal/models",

				ServicesFolder: "internal/services",

				MigrationsFolder: "internal/migrations",

				RepositoryFolder: "internal/repository",
				DatabaseFolder:   "internal/database",
			}

			return configPaths, nil

		case "Fiber":

		case "Echo":

		}

	case "Python":
		switch framework {
		case "FastApi":

		case "Django":

		case "Flask":

		}

	case "JavaScript":
		switch framework {
		case "Express":

		case "NestJS":

		case "Fastify":
		}

	case "TypeScript":
		switch framework {
		case "Express":

		case "NestJS":

		case "Fastify":
		}
	}

	return configPaths, nil
}
