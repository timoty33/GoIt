package create

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"os"
	"strings"
)

const placeholder = "// goit:add-routes-here"

// createRoute gera a string da rota a ser inserida no arquivo.
func createRoute(routeName, method, handlerName string) (string, error) {
	return fmt.Sprintf(`r.%s("/%s", handler.%s)`, strings.ToUpper(method), routeName, handlerName), nil
}

// UpdateRoutesFile lê o arquivo de rotas, insere a nova rota e salva o arquivo.
func UpdateRoutesFile(routeName, method, handlerName string, configs utils.ConfigPaths) error {
	newRouteLine, err := createRoute(routeName, method, handlerName)
	if err != nil {
		return err
	}

	content, err := file.ReadFile(configs.RoutesFile)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler arquivo: %w", err)
	}

	// Usa a função utilitária para inserir o conteúdo. A indentação é adicionada aqui.
	newContent, err := utils.InsertAfterPlaceholder(content, placeholder, "\t"+newRouteLine)
	if err != nil {
		return fmt.Errorf("❌ Falha ao atualizar o arquivo de rotas '%s': %w", configs.RoutesFile, err)
	}

	return os.WriteFile(configs.RoutesFile, []byte(newContent), 0644)
}
