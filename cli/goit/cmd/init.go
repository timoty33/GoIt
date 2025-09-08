package cmd

import (
	"fmt"
	"goit/cli/goit/cmd/structure"
	"goit/cli/goit/cmd/structure/setup"
	"goit/utils"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func formulario() (string, string, string, string, string, bool, string) {
	var nomeProjeto string
	var tipoProjeto string
	var linguagemProjeto string
	var frameworkProjeto string
	var dbProjeto string
	var otherTemplate string
	var templatePath string

	// Pergunta 0: template alternativo
	survey.AskOne(&survey.Select{
		Message: "Do you want to use a customized template: ",
		Options: []string{"Yes", "No"},
		Default: "No",
	}, &otherTemplate)

	if otherTemplate == "Yes" {
		survey.AskOne(&survey.Input{
			Message: "Write the path to template [path/to/my/template/files.tmpl]: ",
		}, &templatePath)

		// Pergunta 1: nome do projeto
		survey.AskOne(&survey.Input{
			Message: "What is the name of the project: ",
		}, &nomeProjeto)

		return nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto, true, templatePath
	}

	// Pergunta 1: nome do projeto
	survey.AskOne(&survey.Input{
		Message: "What is the name of the project: ",
	}, &nomeProjeto)

	// Pergunta 2: tipo do projeto (escolha única)
	survey.AskOne(&survey.Select{
		Message: "What is the type of the project: ",
		Options: []string{"Frontend", "Backend", "FullStack"},
		Default: "Backend",
	}, &tipoProjeto)

	// Pergunta 3: linguagens do projeto
	if tipoProjeto == "FullStack" || tipoProjeto == "Backend" {
		survey.AskOne(&survey.Select{
			Message: "What is the language of the project: ",
			Options: []string{"Go", "JavaScript", "TypeScript"},
		}, &linguagemProjeto)
	}

	// Pergunta 4: Framework do projeto
	switch linguagemProjeto {
	case "Go":

		survey.AskOne(&survey.Select{
			Message: "What is the framework: ",
			Options: []string{"Gin", "Echo", "Fiber"},
		}, &frameworkProjeto)

	case "JavaScript", "TypeScript":

		survey.AskOne(&survey.Select{
			Message: "What is the framework: ",
			Options: []string{"Express", "NestJS", "Fastify"},
		}, &frameworkProjeto)

	}

	// Pergunta 5: DB do projeto
	survey.AskOne(&survey.Select{
		Message: "What is the DB of project: ",
		Options: []string{"SQLite", "PostgreSQL", "MySQL", "MongoDB", "None"},
	}, &dbProjeto)

	return nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto, false, ""
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa um novo projeto GoIt",
	Long: `Cria uma estrutura de diretórios e arquivos padrão para um novo projeto (backend, frontend, fullstack)

Exemplo:
  goit init`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var orm string

		nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto, otherTemplate, templatePath := formulario()

		projectPath := filepath.Join(nomeProjeto)
		fmt.Println("Creating project...")

		// utilizando outro template
		if otherTemplate {
			err := structure.CreateStructureOther(templatePath, projectPath)
			if err != nil {
				return fmt.Errorf("erro ao carregar template: %w", err)
			}
		} else {
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
				Ignore:             []string{"*.md", "*.txt", "./bin/**/*"},
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

			configPaths, err := structure.CreateStructure(projectPath, linguagemProjeto, frameworkProjeto, tipoProjeto, dbProjeto)
			if err != nil {
				return fmt.Errorf("error creating structure: %w", err)
			}

			err = os.MkdirAll(filepath.Join(projectPath, ".goit", "config"), 0755)
			if err != nil {
				return fmt.Errorf("error creating .goit/config directory: %w", err)
			}
			err = utils.SaveJsonConfigs(configs, configPaths, projectPath)
			if err != nil {
				return fmt.Errorf("error creating configs: %w", err)
			}

		}

		fmt.Println("Project created!")

		fmt.Println("Setuping project...")

		fmt.Print(`Do you want to install dependencies:
[1] Yes
[2] No
>> `)
		var install string
		fmt.Scanln(&install)

		// Se o usuário escolher instalar as dependências
		if install == "1" {
			switch linguagemProjeto {
			case "Go":
				if err := setup.GoModInit(projectPath, projectPath); err != nil {
					return fmt.Errorf("error setuping go.mod: %w", err)
				}

				err := setup.InstallDependenciesGo(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("error installing dependencies: %w", err)
				}

				if tipoProjeto == "Frontend" || tipoProjeto == "FullStack" {
					// instala dependencias do front
					utils.CmdExecuteInDir(filepath.Join(nomeProjeto, "frontend"), "npm", "i")
				}

			case "JavaScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("error setuping Node: %w", err)
				}

				err := setup.InstallDependenciesJS(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("error installing dependencies: %w", err)
				}
				if tipoProjeto == "Frontend" || tipoProjeto == "FullStack" {
					// instala dependencias do front
					utils.CmdExecuteInDir(filepath.Join(nomeProjeto, "frontend"), "npm", "i")
				}

			case "TypeScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("error setuping Node: %w", err)
				}

				if err := setup.TsInit(projectPath); err != nil {
					return fmt.Errorf("error setuping TypeScript: %w", err)
				}

				err := setup.InstallDependenciesTS(projectPath, frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("error installing dependedencies: %w", err)
				}

				if tipoProjeto == "Frontend" || tipoProjeto == "FullStack" {
					// instala dependencias do front
					utils.CmdExecuteInDir(filepath.Join(nomeProjeto, "frontend"), "npm", "i")
				}
			}
		} else {
			switch linguagemProjeto {
			case "Go":
				if err := setup.GoModInit(projectPath, projectPath); err != nil {
					return fmt.Errorf("error setuping go.mod: %w", err)
				}

			case "JavaScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("error setuping Node: %w", err)
				}

			case "TypeScript":
				if err := setup.NodeInit(projectPath); err != nil {
					return fmt.Errorf("error setuping Node: %w", err)
				}

				if err := setup.TsInit(projectPath); err != nil {
					return fmt.Errorf("error setuping Typescript: %w", err)
				}
			}
		}

		fmt.Println("Setup completed!")

		fmt.Println("Enjoy ;)")

		return nil
	},
}
