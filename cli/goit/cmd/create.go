package cmd

import (
	"fmt"
	"goit/cli/goit/cmd/create"
	"goit/utils"
	"goit/utils/file"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var createHandler = &cobra.Command{
	Use:   "handler [name] [flags]",
	Short: "Create a Handler and a route with it",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		nome := args[0]

		switch configProject.ProgrammingLanguage {
		case "Go":
			switch configProject.Framework {
			case "Gin":

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

				return nil
			}

		case "TypeScript":

		case "JavaScript":

		}

		return nil
	},
}

var createMiddleware = &cobra.Command{
	Use:   "middleware [name] [flags]",
	Short: "Cria um midleware e atribui ele no server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		nome := args[0]

		switch configProject.ProgrammingLanguage {
		case "Go":
			switch configProject.Framework {
			case "Gin":
				nameVerify, err := utils.TitleNameVerify(nome)
				if err != nil {
					return fmt.Errorf("❌ The name can't be used: %w", err)
				}

				if newFile {
					err := create.CreateMiddlewareNewFile(configsPath, nome, nameVerify)
					if err != nil {
						return fmt.Errorf("❌ Error creating middleware: %w", err)
					}
				} else {
					err := create.CreateMiddleware(nameVerify, configsPath)
					if err != nil {
						return fmt.Errorf("❌ Error injecting middleware: %w", err)
					}
				}

				fmt.Println("Middleware created!")

				return nil
			}
		}

		return nil
	},
}

var createRoute = &cobra.Command{
	Use:   "route [name] [flags]",
	Short: "Cria uma rota atribuindo um handler",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		nome := args[0]

		switch configProject.ProgrammingLanguage {
		case "Go":
			switch configProject.Framework {
			case "Gin":
				if handlerName == "" {
					return fmt.Errorf("❌ Você não definiu o nome de um handler para ser atribuído a rota, %s", handlerName)
				}

				err = create.UpdateRoutesFile(nome, method, handlerName, configsPath)
				if err != nil {
					return fmt.Errorf("❌ Erro ao criar rota, %w", err)
				}

				fmt.Printf("Rota %s criada com sucesso", nome)

				return nil
			}

		case "TypeScript":

		case "JavaScript":

		default:
			fmt.Println("Programming language not suported...", configProject.ProgrammingLanguage)
			return fmt.Errorf("programming language not suported", configProject.ProgrammingLanguage)
		}

		return nil
	},
}

var createEnv = &cobra.Command{
	Use:   "env [file_name]",
	Short: "Create a .env file and a file loads the .env and have a get with fallback function",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configProject, _, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		nome := args[0]

		switch configProject.ProgrammingLanguage {
		case "Go":
			content := `package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadDotEnv() (error) {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("erro ao obter diretório absoluto: %w", err)
	}
	
	envFilePath := filepath.Join(currentDir, "..", "env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		return fmt.Errorf("erro ao carregar env: %w", err)
	}
	
	return nil
}

// função helper para pegar env vars
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
`

			env, err := os.Create(".env")
			if err != nil {
				return fmt.Errorf("❌ Error creating .env file: %w", err)
			}
			defer env.Close()

			file.CreateArqVerify("internal/config", filepath.Join("internal", "config", nome+".go"), nome+".go", content)

			var install string
			fmt.Print(`Do you want to install godotenv
[1] Yes
[2] No
>> `)
			fmt.Scanln(&install)

			if install == "1" {
				cmd := exec.Command("go", "get", "github.com/joho/godotenv")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

				fmt.Println("godotenv installed with success!")
			}

			fmt.Println("Env criado com sucesso!")

			return nil
		}

		return nil
	},
}

var createDto = &cobra.Command{
	Use:   "dto [name] [flags]",
	Short: "Create a dto and insert in the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		nome := args[0]

		switch configProject.ProgrammingLanguage {
		case "Go":
			if handlerName == "" {
				return fmt.Errorf("the handler name is required, try --for or -H")
			}

			if strings.ToLower(dtoMode) != "input" && dtoMode != "output" {
				return fmt.Errorf("invalid value for --dto-mode. Use 'input' or 'output'")
			}

			nameVerify, err := utils.TitleNameVerify(nome)
			if err != nil {
				return fmt.Errorf("Name cannot be used: %w", err)
			}

			if newFile {
				err := create.CreateDtoNewFile(camps, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("Error creating new DTO: %w", err)
				}
			} else {
				err := create.CreateDto(camps, nameVerify, configsPath)
				if err != nil {
					return fmt.Errorf("Error creating inline DTO: %w", err)
				}
			}

			err = create.UpdateHandlerWithDto(dtoMode, handlerName, nameVerify, configsPath)
			if err != nil {
				return fmt.Errorf("Error injecting DTO into handler: %w", err)
			}

			fmt.Println("DTO created and handler updated successfully!")
		}

		return nil
	},
}

var createCmd = &cobra.Command{
	Use:   "create [tipo] [nome]",
	Short: "Create code and files for you with the configs in .goit/config/goit.config.json",
	Args:  cobra.ExactArgs(2), // tipo de criação e nome
	RunE: func(cmd *cobra.Command, args []string) error {
		tipo := args[0]
		nome := args[1]

		configsProject, configsPath, err := utils.LoadJsonConfig()
		if err != nil {
			return fmt.Errorf("❌ Erro: %w", err)
		}

		switch tipo {

		case "env":
		case "route":
		case "dto":

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
				fmt.Print(`Do you want to install GORM:
[1] Yes
[2] No
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
		}

		return nil
	},
}
