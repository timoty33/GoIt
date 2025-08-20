package cmd

import (
	"fmt"
	"goit/goit/cmd/structure"
	"goit/goit/cmd/structure/setup"
	"goit/utils"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func formulario() (string, string, string, string, string) {
	var nomeProjeto string
	var tipoProjeto string
	var linguagemProjeto string
	var frameworkProjeto string
	var dbProjeto string

	// Pergunta 1: nome do projeto
	survey.AskOne(&survey.Input{
		Message: "Qual o nome do projeto: ",
	}, &nomeProjeto)

	// Pergunta 2: tipo do projeto (escolha única)
	survey.AskOne(&survey.Select{
		Message: "Escolha o tipo de projeto: Frontend e FullStack ainda em desenvolvimento, sem suporte, não use",
		Options: []string{"Frontend", "Backend", "FullStack"},
		Default: "Backend",
	}, &tipoProjeto)

	// Pergunta 3: linguagens do projeto
	if tipoProjeto == "FullStack" || tipoProjeto == "Backend" {
		survey.AskOne(&survey.Select{
			Message: "Escolha a linguagem do backend: (você precisa ter a linguagem instalada no seu computador)",
			Options: []string{"Go", "JavaScript", "TypeScript"},
		}, &linguagemProjeto)
	}

	// Pergunta 4: Framework do projeto
	switch linguagemProjeto {
	case "Go":

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"Gin", "Echo", "Fiber"},
		}, &frameworkProjeto)

	case "JavaScript", "TypeScript":

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"Express", "NestJS", "Fastify"},
		}, &frameworkProjeto)

	}

	// Pergunta 5: DB do projeto
	survey.AskOne(&survey.Select{
		Message: "Escolha o banco de dados que será usado: MongoDB ainda em desenvolvimento, sem suporte, não use",
		Options: []string{"SQLite", "PostgreSQL", "MySQL", "MongoDB", "Nenhuma"},
	}, &dbProjeto)

	return nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa um novo projeto GoIt",
	Long: `Cria uma estrutura de diretórios e arquivos padrão para um novo projeto (backend, frontend, fullstack)

Exemplo:
  goit init`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var orm string

		nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto := formulario()

		projectPath := filepath.Join(nomeProjeto)
		fmt.Println("Iniciando o projeto:", projectPath)

		var initCommandBackend string
		switch linguagemProjeto {
		case "Go":
			orm = "gorm"
			initCommandBackend = "go run main.go"
		case "JavaScript":
			orm = "prisma"
			initCommandBackend = "node src/index.js"
		case "TypeScript":
			orm = "prisma"
			initCommandBackend = "npx ts-node src/index.ts"
		}

		lintConfig := utils.LintType{
			Lint:         true,
			LintApply:    false,
			Format:       true,
			LintFrontEnd: "frontend/",
			LintBackEnd:  ".",
		}
		// structs for config
		hotBack := utils.HotreloadBackend{
			Active:     true,
			ListenPath: ".",
		}
		hotFront := utils.HotreloadFrontend{
			Active:     true,
			ListenPath: "./frontend",
		}
		devConfig := utils.Dev{
			HotReloadBackend:   hotBack,
			HotReloadFrontend:  hotFront,
			Ignore:             []string{"md", "*.txt", "./bin/**/*"},
			InitCommandBackend: initCommandBackend,
		}
		configRun := utils.Run{
			Lint: lintConfig,
			Dev:  devConfig,
		}

		configs := utils.ConfigProject{
			ProjectName:         projectPath,
			ProjectType:         tipoProjeto,
			ProgrammingLanguage: linguagemProjeto,
			Framework:           frameworkProjeto,
			DataBase:            dbProjeto,
			Port:                "8080",
			Orm:                 orm,
			Run:                 configRun,
		}

		configPaths, err := structure.CreateStructure(projectPath, linguagemProjeto, frameworkProjeto, tipoProjeto)
		if err != nil {
			return fmt.Errorf("erro ao criar estrutura: %w", err)
		}

		err = utils.SaveJsonConfigs(configs, configPaths, projectPath)
		if err != nil {
			return fmt.Errorf("erro ao salvar configurações: %w", err)
		}

		fmt.Println("Estrutura do projeto criada com sucesso!")

		fmt.Println("Iniciando o setup do projeto...")

		fmt.Print(`Deseja instalar as dependências?
[1] Sim
[2] Não
>> `)
		var install string
		fmt.Scanln(&install)

		// Se o usuário escolher instalar as dependências
		if install == "1" {
			switch linguagemProjeto {
			case "Go":
				if err := setup.GoModInit(projectPath, projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o módulo Go: %w", err)
				}

				err = setup.InstallDependenciesGo(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}

			case "JavaScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

				err = setup.InstallDependenciesJS(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}

			case "TypeScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

				if err = setup.TsInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o TypeScript: %w", err)
				}

				err = setup.InstallDependenciesTS(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}
			}
		} else {
			fmt.Println("Você escolheu não instalar as dependências.")

			switch linguagemProjeto {
			case "Go":
				if err := setup.GoModInit(projectPath, projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o módulo Go: %w", err)
				}

			case "JavaScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

			case "TypeScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

				if err = setup.TsInit(projectPath); err != nil {
					return fmt.Errorf("erro ao inicializar o TypeScript: %w", err)
				}
			}
		}

		fmt.Println("Setup do projeto concluído com sucesso!")

		return nil
	},
}
