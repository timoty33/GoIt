package create

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"os"
	"path/filepath"
	"strings"
)

func CreateMiddlewareNewFile(configPaths utils.ConfigPaths, nomeMiddleware, nameVerify string) error {
	// montar o content:
	content := fmt.Sprintf(`package middleware

import (
	"github.com/gin-gonic/gin"
)
	
%s`, createMiddlewareFuncString(nameVerify))

	// criar o arquivo
	err := file.CreateArqVerify(configPaths.MiddlewaresFolder, filepath.Join(configPaths.MiddlewaresFolder, nomeMiddleware+".go"), nomeMiddleware+".go", content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao criar arquivo: %w", err)
	}

	return InjectMiddlewareUse(configPaths, nameVerify)
}

func CreateMiddleware(nameVerify string, configsPath utils.ConfigPaths) error {

	// ler o arquivo de middleware
	middlewareFileContent, err := file.ReadFile(configsPath.MiddlewaresFile)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler arquivo de middleware: %w", err)
	}

	// injetar o middleware no arquivo
	newContent, err := utils.InsertAfterPlaceholder(middlewareFileContent, "// goit:add-middlewares-here", createMiddlewareFuncString(nameVerify))
	if err != nil {
		return fmt.Errorf("❌ Erro ao injetar middleware no arquivo de middleware: %w", err)
	}

	err = os.WriteFile(configsPath.MiddlewaresFile, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("❌ Erro ao escrever o arquivo de middleware: %w", err)
	}

	return InjectMiddlewareUse(configsPath, nameVerify)
}

func createMiddlewareFuncString(nomeMiddleware string) string {
	return fmt.Sprintf(`func %s() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}	
}`, nomeMiddleware)
}

func InjectMiddlewareUse(config utils.ConfigPaths, name string) error {
	serverFileContent, err := file.ReadFile(config.ServerFile)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler arquivo de servidor: %w", err)
	}

	if strings.Contains(serverFileContent, fmt.Sprintf("r.Use(%s())", name)) {
		return nil // já foi injetado
	}

	newContent, err := utils.InsertAfterPlaceholder(serverFileContent, "// goit:add-middlewares-here", fmt.Sprintf("\tr.Use(middleware.%s())\n", name))
	if err != nil {
		return fmt.Errorf("❌ Erro ao injetar middleware no arquivo de servidor: %w", err)
	}

	return os.WriteFile(config.ServerFile, []byte(newContent), 0644)
}
