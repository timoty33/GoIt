package cmd

import (
	"fmt"
	"goit/cmd/create"
	"goit/utils"
	"goit/utils/file"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [tipo] [nome]",
	Short: "Cria uma parte de código, ou um arquivo já iniciado, e acrescenta em outros arquivos de acordo com o .goit.config.json",
	Args:  cobra.ExactArgs(2), // tipo de criação e nome
	RunE: func(cmd *cobra.Command, args []string) error {
		tipo := args[0]
		nome := args[1]

		configsProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		switch tipo {
		case "handler":

			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("❌ O nome não pode ser usado: %w", err)
			}

			if newFile {
				err := create.CreateHandlerFile(nome, method, configsPath, handlerParams)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar handler: %w", err)
				}
			} else {
				err := create.CreateHandler(nome, method, configsPath, handlerParams)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar handler: %w", err)
				}
			}

			if routeName == "" {
				err = create.UpdateRoutesFile(nome, method, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao adicionar a rota: %w", err)
				}
			} else {
				err = create.UpdateRoutesFile(routeName, method, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao adicionar a rota: %w", err)
				}
			}

			fmt.Println("Handler criado com sucesso!")

		case "env":
			content := `package config
import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/joho/godotenv"
)
	
func loadDotEnv() (string, error) {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", fmt.Errorf("erro ao obter diretório absoluto: %w", err)
	}
	
	envFilePath := filepath.Join(currentDir, "..", "env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		return "", fmt.Errorf("erro ao carregar env: %w", err)
	}

	apiKey := os.Getenv("API_KEY")
	
	return apiKey, nil
}`

			env, err := os.Create(".env")
			if err != nil {
				return fmt.Errorf("❌ Erro ao criar o arquivo .env: %w", err)
			}
			defer env.Close()

			file.CreateArqVerify("internal/config", filepath.Join("internal", "config", nome+".go"), nome+".go", content)

			var install string
			fmt.Print(`Você quer instalar o GoDotEnv?
[1] Sim
[2] Não
>> `)
			fmt.Scanln(&install)

			if install == "1" {
				cmd := exec.Command("go", "get", "github.com/joho/godotenv")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

				fmt.Println("Instalação completa!")
			}

			fmt.Println("Env criado com sucesso!")

		case "route":
			if handlerName == "" {
				return fmt.Errorf("❌ Você não definiu o nome de um handler para ser atribuído a rota, %s", handlerName)
			}

			err := create.UpdateRoutesFile(nome, method, handlerName, configsPath)
			if err != nil {
				return fmt.Errorf("❌ Erro ao criar rota, %w", err)
			}

			fmt.Printf("Rota %s criada com sucesso", nome)

		case "dto":
			if handlerName == "" {
				return fmt.Errorf("o nome do handler é obrigatório. Use --for ou -H")
			}

			if dtoMode != "input" && dtoMode != "output" {
				return fmt.Errorf("valor inválido para --dto-mode. Use 'input' ou 'output'")
			}

			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("❌ O nome não pode ser usado: %w", err)
			}

			if newFile {
				err := create.CreateDtoNewFile(camps, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar DTO novo: %w", err)
				}
			} else {
				err := create.CreateDto(camps, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar DTO embutido: %w", err)
				}
			}

			err = create.UpdateHandlerWithDto(dtoMode, handlerName, nameVerify, configsPath)
			if err != nil {
				return fmt.Errorf("❌ Erro ao injetar DTO no handler: %w", err)
			}

			fmt.Println("DTO criado e handler atualizado com sucesso!")

		case "model":
			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("❌ O nome não pode ser usado: %w", err)
			}

			create.CreateModelNewFile(nome, nameVerify, configsPath, configsProject, camps)
			if err != nil {
				return fmt.Errorf("❌ Erro ao criar modelo: %w", err)
			}
			fmt.Println("Modelo criado com sucesso!")

			var install string
			if configsProject.Orm == "gorm" {
				fmt.Print(`Você deseja instalar o GORM
[1] Sim
[2] Não
>> `)
				fmt.Scanln(&install)

				if install == "1" {
					cmd := exec.Command("go", "get", "gorm.io/gorm")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Run()

					fmt.Println("Instalação completa!")
				}
			}

		case "migration":
			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("❌ O nome não pode ser usado: %w", err)
			}

			err = create.CreateMigration(nome, nameVerify, modelName, configsPath, configsProject)
			if err != nil {
				return fmt.Errorf("❌ Erro ao criar migration: %w", err)
			}

			fmt.Println("Migration criada com sucesso!")

		case "middleware":
			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("❌ O nome não pode ser usado: %w", err)
			}

			if newFile {
				err := create.CreateMiddlewareNewFile(configsPath, nome, nameVerify)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar middleware: %w", err)
				}
			} else {
				err := create.CreateMiddleware(nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao injetar middleware: %w", err)
				}
			}

			fmt.Println("Middleware criado com sucesso!")

		}

		return nil
	},
}
