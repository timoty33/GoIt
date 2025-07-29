package create

import (
	"fmt"
	"os"
	"path/filepath"

	"goit/utils"
	"goit/utils/file"
)

func CreateHandlerFile(nomeHandler, method string, configs utils.ConfigPaths, params []string) error {
	nomeFunc, err := utils.TitleNameVerify(nomeHandler)
	if err != nil {
		return fmt.Errorf("❌ O nome não pode ser usado, %w", err)
	}

	content := `package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ` + nomeFunc + `(c *gin.Context) {
`

	for _, param := range params {
		content += fmt.Sprintf("\t%s := c.Query(\"%s\")\n", param, param)
	}

	content += fmt.Sprintf("\n\tc.JSON(http.StatusOK, gin.H{\"message\": \"Handler '%s' funcionando!\"})\n", nomeHandler)
	content += "}\n"

	fullPath := filepath.Join(configs.HandlersFolder, nomeHandler+".go")

	if err := file.CreateArqVerify(configs.HandlersFolder, fullPath, nomeHandler, content); err != nil {
		return fmt.Errorf("❌ Algum erro aconteceu ao criar o arquivo: %w", err)
	}

	return nil
}

func CreateHandler(nomeHandler, method string, configs utils.ConfigPaths, params []string) error {
	nomeFunc, err := utils.TitleNameVerify(nomeHandler)
	if err != nil {
		return fmt.Errorf("❌ O nome não pode ser usado, %w", err)
	}

	content := fmt.Sprintf("func %s(c *gin.Context) {\n", nomeFunc)

	for _, param := range params {
		content += fmt.Sprintf("\t%s := c.Query(\"%s\")\n", param, param)
	}

	content += fmt.Sprintf("\n\tc.JSON(http.StatusOK, gin.H{\"message\": \"Handler '%s' funcionando!\"})\n", nomeHandler)
	content += "}\n\n"

	handleFileContentByte, err := os.ReadFile(configs.HandlersFile)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler o arquivo de handlers, %w", err)
	}

	handleFileContent := string(handleFileContentByte)

	newContentFile, err := utils.InsertAfterPlaceholder(handleFileContent, "goit:add-handlers-here", content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao injetar handler no arquivo, %w", err)
	}

	return os.WriteFile(configs.HandlersFile, []byte(newContentFile), 0644)
}
