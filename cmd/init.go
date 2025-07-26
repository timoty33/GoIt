package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"goit/cmd/structure"

	"github.com/spf13/cobra"
)

func getModGin(projectPath string) error {
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
	Use:   "init [nome-do-projeto]",
	Short: "Inicializa um novo projeto GoIt",
	Long: `Cria uma estrutura de diretórios e arquivos padrão para um novo projeto Go.

Exemplo:
  goit init meu-projeto-incrivel -f gin --install`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprint(os.Stderr, "O nome do projeto é obrigatório")
			os.Exit(1)
		}

		nomeProjeto := args[0]

		fmt.Printf("🚀 Iniciando a criação do projeto '%s' com o framework '%s'...\n", nomeProjeto, framework)

		if err := createEstructure(nomeProjeto, framework); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Erro ao criar a estrutura do projeto: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Estrutura de diretórios e arquivos criada com sucesso.")

		if err := structure.GoModInit(nomeProjeto); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Erro ao inicializar o go.mod: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ go.mod inicializado.")

		if install {
			fmt.Println("📦 Instalando dependências...")
			if err := getModGin(nomeProjeto); err != nil {
				fmt.Fprintf(os.Stderr, "❌ Erro ao instalar dependências: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Dependências instaladas.")
		}

		fmt.Printf("\n🎉 Projeto '%s' criado com sucesso!\n\nPara começar:\n  cd %s\n  go run cmd/main.go\n", nomeProjeto, nomeProjeto)
	},
}
