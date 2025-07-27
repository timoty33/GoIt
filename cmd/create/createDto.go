package create

import (
	"fmt"
	"goit/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CreateDto(camps []string, dtoName string, configs utils.Config) error {
	var content2 string

	content1 := fmt.Sprintf("type %s struct{\n", dtoName)
	for _, param := range camps {
		content2 += fmt.Sprintf("\t%s\n", param)
	}
	content := content1 + content2 + "}"

	fileContent, err := utils.ReadFile(configs.DtoFile)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler arquivo, %w", err)
	}

	newFileContent, err := utils.InsertAfterPlaceholder(fileContent, "goit:add-dtos-here", content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao usar marcador, %w", err)
	}

	err = os.WriteFile(configs.DtoFile, []byte(newFileContent), 0644)
	if err != nil {
		return fmt.Errorf("❌ Erro ao injetar código, %w", err)
	}

	return nil
}

func CreateDtoNewFile(camps []string, dtoName string, configs utils.Config) error {
	// Define nome do arquivo (ex: userinput.go) e caminhos
	fileName := strings.ToLower(dtoName) + ".go"
	fullPath := filepath.Join(configs.DtoFolder, fileName)
	folder := configs.DtoFolder

	// Monta os campos do struct
	var fields string
	for _, param := range camps {
		fields += fmt.Sprintf("\t%s\n", param)
	}

	// Conteúdo do arquivo DTO
	content := fmt.Sprintf(`package %s

type %s struct {
%s}
`, filepath.Base(folder), dtoName, fields)

	// Usa função utilitária para criar e salvar o arquivo com verificação
	err := utils.CreateArqVerify(folder, fullPath, fileName, content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao criar DTO: %w", err)
	}

	fmt.Printf("✅ DTO %s criado com sucesso em: %s\n", dtoName, fullPath)
	return nil
}

func UpdateHandlerWithDto(mode, fullPath, handlerName string) error {
	content, err := utils.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler arquivo: %w", err)
	}

	// Regex para encontrar a função do handler
	caso := fmt.Sprintf(`func %s\(c \*gin.Context\) \{`, handlerName)
	re, err := regexp.Compile(caso)
	if err != nil {
		return fmt.Errorf("❌ Erro ao utilizar regex: %w", err)
	}

	// Criar o código que será injetado, baseado no modo
	var dtoCode string
	mode = strings.ToUpper(mode)
	switch mode {
	case "INPUT":
		dtoCode = createDtoContentInput(handlerName + "Input")
	case "OUTPUT":
		dtoCode = createDtoContentOutput(handlerName + "Output")
	default:
		return fmt.Errorf("❌ Mode inválido: %s", mode)
	}

	// Injeta o código após a abertura da função
	newContent := re.ReplaceAllStringFunc(content, func(match string) string {
		return match + "\n" + dtoCode
	})

	// Escreve de volta no arquivo
	if err := os.WriteFile(fullPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("❌ Erro ao salvar alterações: %w", err)
	}

	fmt.Println("✅ DTO injetado com sucesso no handler.")
	return nil
}

func createDtoContentInput(dtoName string) string {
	return fmt.Sprintf(`	var input %s
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}`, dtoName)
}

func createDtoContentOutput(dtoName string) string {
	return fmt.Sprintf(`	output := %s{
		// preencha os campos aqui
	}
	c.JSON(200, output)`, dtoName)
}
