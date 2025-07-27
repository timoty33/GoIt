package create

import (
	"fmt"
	"os"
	"path/filepath"

	"goit/utils"
)

func CreateHandlerFile(nomeHandler, method string, configs utils.Config) error {
	// Validação de nome (evita espaços, símbolos e inicia com número)
	nomeFunc, err := utils.TitleNameVerify(nomeHandler)
	if err != nil {
		return fmt.Errorf("❌ O nome não pode ser usado, %w", err)
	}

	// Conteúdo do handler gerado
	content := `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ` + nomeFunc + `(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handler '` + nomeHandler + `' funcionando!"})
}
`

	fullPath := filepath.Join(configs.HandlersFolder, nomeHandler+".go")

	err = utils.CreateArqVerify(configs.HandlersFolder, fullPath, nomeHandler, content)
	if err != nil {
		return fmt.Errorf("❌ Algum erro aconteceu ao criar o arquivo: %w", err)
	}

	return nil
}

func CreateHandler(nomeHandler, method string, configs utils.Config) error {
	nomeFunc, err := utils.TitleNameVerify(nomeHandler)
	if err != nil {
		return fmt.Errorf("❌ O nome não pode ser usado, %w", err)
	}

	content := `func ` + nomeFunc + `(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Handler '` + nomeHandler + `' funcionando!"})
}
`
	handleFilePath := filepath.Join(configs.HandlersFile)
	handleFileContentByte, err := os.ReadFile(handleFilePath)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler o arquivo de handlers, %w", err)
	}
	handleFileContent := string(handleFileContentByte)

	newContentFile, err := utils.InsertAfterPlaceholder(handleFileContent, "goit:add-handlers-here", content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao injetar handler no arquivo, %w", err)
	}

	return os.WriteFile(handleFilePath, []byte(newContentFile), 0644)
}
