package cmd

import (
	"fmt"
	"goit/cmd/structure"
	"goit/cmd/structure/setup"
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
			Options: []string{"Go", "JavaScript", "TypeScript", "Python"},
		}, &linguagemProjeto)
	}

	// Pergunta 4: Framework do projeto
	switch linguagemProjeto {
	case "Go":

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"Gin", "Echo", "Fiber"},
		}, &frameworkProjeto)

	case "Python":

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"FastApi", "Django", "Flask"},
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

		fmt.Println("Iniciando o projeto:", nomeProjeto)

		switch linguagemProjeto {
		case "Go":
			orm = "gorm"
		case "Python":
			orm = "sqlalchemy"
		case "JavaScript", "TypeScript":
			orm = "prisma"
		}

		configs := utils.ConfigProject{
			ProjectName:         nomeProjeto,
			ProgrammingLanguage: linguagemProjeto,
			Framework:           frameworkProjeto,
			DataBase:            dbProjeto,
			Port:                "8080",
			Orm:                 orm,
			HotReload:           true,
		}

		configPaths, err := structure.CreateStructure(nomeProjeto, linguagemProjeto, frameworkProjeto, tipoProjeto)
		if err != nil {
			return fmt.Errorf("erro ao criar estrutura: %w", err)
		}

		err = utils.SaveJsonConfigs(configs, configPaths, filepath.Join(nomeProjeto))
		if err != nil {
			return fmt.Errorf("erro ao salvar configurações: %w", err)
		}

		fmt.Println("Estrutura do projeto criada com sucesso!")

		fmt.Println("Iniciando o setup do projeto...")

		fmt.Print(`Deseja instalar as dependências?
[1] Sim
[2] Não
>>`)
		var install string
		fmt.Scanln(&install)

		// Se o usuário escolher instalar as dependências
		if install == "1" {
			switch linguagemProjeto {
			case "Go":
				if err := setup.GoModInit(nomeProjeto); err != nil {
					return fmt.Errorf("erro ao inicializar o módulo Go: %w", err)
				}

				err = setup.InstallDependenciesGo(frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}

			case "Python":
				if err := setup.PythonInit(nomeProjeto); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Python (venv): %w", err)
				}

				err = setup.InstallDependenciesPython(frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}

			case "JavaScript":
				if err := setup.NodeInit(nomeProjeto); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

				err = setup.InstallDependenciesJS(frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}

			case "TypeScript":
				if err := setup.NodeInit(nomeProjeto); err != nil {
					return fmt.Errorf("erro ao inicializar o projeto Node.js: %w", err)
				}

				if err = setup.TsInit(nomeProjeto); err != nil {
					return fmt.Errorf("erro ao inicializar o TypeScript: %w", err)
				}

				err = setup.InstallDependenciesTS(frameworkProjeto, dbProjeto)
				if err != nil {
					return fmt.Errorf("erro ao instalar as dependências: %w", err)
				}
			}
		}

		fmt.Println("Setup do projeto concluído com sucesso!")

		return nil
	},
}
