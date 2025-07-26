package create

import (
	"fmt"
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
