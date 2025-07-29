package cmd

import (
	"fmt"

	"goit/cmd/structure"

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
		Message: "Qual o nome do projeto?",
	}, &nomeProjeto)

	// Pergunta 2: tipo do projeto (escolha única)
	survey.AskOne(&survey.Select{
		Message: "Escolha o tipo de projeto:",
		Options: []string{"Frontend", "Backend", "FullStack"},
		Default: "Backend",
	}, &tipoProjeto)

	// Pergunta 3: linguagens do projeto
	if tipoProjeto == "FullStack" || tipoProjeto == "Backend" {

		survey.AskOne(&survey.Select{
			Message: "Escolha a linguagem do backend:",
			Options: []string{"Go", "JavaScript", "TypeScript", "Python"},
		}, &linguagemProjeto)
	}

	// Pergunta 4: Framework do projeto
	if linguagemProjeto == "Go" {

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"Gin", "Echo", "Fiber"},
		}, &frameworkProjeto)

	} else if linguagemProjeto == "Python" {

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"FastAPI", "Django", "Flask"},
		}, &frameworkProjeto)

	} else if linguagemProjeto == "JavaScript" || linguagemProjeto == "TypeScript" {

		survey.AskOne(&survey.Select{
			Message: "Escolha o framework que será usado:",
			Options: []string{"Express", "NestJS", "Fastify"},
		}, &frameworkProjeto)

	}

	// Pergunta 5: DB do projeto
	survey.AskOne(&survey.Select{
		Message: "Escolha o banco de dados que será usado:",
		Options: []string{"SQLite", "PostgreSQL", "MySQL", "MongoDB", "Nenhuma"},
	}, &dbProjeto)

	return nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto
}

func createEstructure(nomeProjeto, framework string) error {
	if framework == "gin" {
		if err := structure.StructureGin(nomeProjeto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("framework '%s' não suportado", framework)
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa um novo projeto GoIt",
	Long: `Cria uma estrutura de diretórios e arquivos padrão para um novo projeto (backend, frontend, fullstack)

Exemplo:
  goit init`,
	Run: func(cmd *cobra.Command, args []string) {
		nomeProjeto, tipoProjeto, linguagemProjeto, frameworkProjeto, dbProjeto := formulario()
	},
}
