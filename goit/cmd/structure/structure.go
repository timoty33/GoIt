package structure

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"os"
	"path/filepath"
)

const (
	permPasta = 0755
)

func CreateStructure(nomeProjeto, linguagem, framework, tipoProjeto string) (utils.ConfigPaths, error) {
	var configPaths utils.ConfigPaths

	err := os.Mkdir(nomeProjeto, permPasta)
	if err != nil {
		return configPaths, fmt.Errorf("erro ao criar pasta do projeto %s: %w", nomeProjeto, err)
	}

	if tipoProjeto == "Backend" || tipoProjeto == "FullStack" {

		// Renderiza os templates (backend, o front será criado com Vite)
		templates, err := file.PercorrerDiretorio(filepath.Join(GetTemplatesPath(), linguagem, framework))
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
				ServerFile: "internal/server/server.go",

				RoutesFile: "internal/routes/routes.go",

				HandlersFile:   "internal/handler/handler.go",
				HandlersFolder: "internal/handler",

				MiddlewaresFolder: "internal/middleware",
				MiddlewaresFile:   "internal/middleware/middleware.go",

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
