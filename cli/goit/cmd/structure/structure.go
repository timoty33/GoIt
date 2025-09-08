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

func CreateStructureOther(templatePath, nomeProjeto string) error {
	err := os.Mkdir(nomeProjeto, permPasta)
	if err != nil {
		return fmt.Errorf("erro ao criar pasta do projeto %s: %w", nomeProjeto, err)
	}

	templates, err := file.PercorrerDiretorio(templatePath)
	if err != nil {
		return fmt.Errorf("erro ao percorrer pasta do diretório\npasta: %s -- %w", templatePath, err)
	}

	// renderizando templates
	RenderTemplates(templates, TemplateData{ProjectName: nomeProjeto}, nomeProjeto)

	return nil
}

func CreateStructure(nomeProjeto, linguagem, framework, tipoProjeto, dbProjeto string) (utils.ConfigPaths, error) {
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

	if tipoProjeto == "Frontend" || tipoProjeto == "FullStack" {
		// rodar com o vite
		fmt.Println("Criando Frontend do projeto...")
		utils.CmdExecuteInDir(filepath.Join(nomeProjeto), "npm", "create", "vite@latest", "frontend")
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

				ModelsFolder: "internal/database/models",

				MigrationsFolder: "internal/database/migrations",

				RepositoryFolder: "internal/database/repositories",
				DatabaseFolder:   "internal/database",
			}

		case "Fiber":

		case "Echo":

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

	// criando .env e env_config para carregar credencias do DB
	if dbProjeto != "None" {
		if dbProjeto != "SQLite" {
			err := utils.CmdExecuteInDir(filepath.Join(nomeProjeto), "goit", "env", "env_config.go")
			if err != nil {
				return configPaths, fmt.Errorf("error creating .env: %w", err)
			}
		}

		templates, err := file.PercorrerDiretorio(filepath.Join(GetTemplatesPath(), linguagem, "db", dbProjeto))
		if err != nil {
			return configPaths, fmt.Errorf("error getting templates: %w", err)
		}

		err = RenderTemplates(templates, TemplateData{ProjectName: nomeProjeto}, filepath.Join(nomeProjeto, configPaths.DatabaseFolder))
		if err != nil {
			return configPaths, fmt.Errorf("error creating db client: %w", err)
		}
	}

	return configPaths, nil
}
