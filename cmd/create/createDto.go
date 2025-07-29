package create

import (
	"fmt"
	"goit/cmd/file"
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

	fileContent, err := file.ReadFile(configs.DtoFile)
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
	err := file.CreateArqVerify(folder, fullPath, fileName, content)
	if err != nil {
		return fmt.Errorf("❌ Erro ao criar DTO: %w", err)
	}

	fmt.Printf("✅ DTO %s criado com sucesso em: %s\n", dtoName, fullPath)
	return nil
}

func UpdateHandlerWithDto(mode, handlerName, dtoName string, configs utils.Config) error {
	handlerFolder := configs.HandlersFolder
	files, err := os.ReadDir(handlerFolder)
	if err != nil {
		return fmt.Errorf("❌ Erro ao ler pasta de handlers: %w", err)
	}

	// Regex para encontrar a função do handler
	caso := fmt.Sprintf(`func %s\(c \*gin.Context\) \{`, handlerName)
	re, err := regexp.Compile(caso)
	if err != nil {
		return fmt.Errorf("❌ Erro ao compilar regex: %w", err)
	}

	// Criar o código que será injetado
	var dtoCode string
	switch strings.ToUpper(mode) {
	case "INPUT":
		dtoCode = createDtoContentInput(dtoName)
	case "OUTPUT":
		dtoCode = createDtoContentOutput(dtoName)
	default:
		return fmt.Errorf("❌ Mode inválido: %s", mode)
	}

	found := false

	for _, arquivo := range files {
		if arquivo.IsDir() || !strings.HasSuffix(arquivo.Name(), ".go") {
			continue
		}

		fullPath := filepath.Join(handlerFolder, arquivo.Name())
		content, err := file.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("❌ Erro ao ler %s: %w", fullPath, err)
		}

		if re.MatchString(content) {
			found = true

			// Injeta o código após abertura da função
			newContent := re.ReplaceAllStringFunc(content, func(match string) string {
				return match + "\n" + dtoCode
			})

			err := os.WriteFile(fullPath, []byte(newContent), 0644)
			if err != nil {
				return fmt.Errorf("❌ Erro ao salvar %s: %w", fullPath, err)
			}

			fmt.Printf("✅ DTO injetado com sucesso no handler %s (%s)\n", handlerName, arquivo.Name())
			break
		}
	}

	if !found {
		return fmt.Errorf("❌ Handler '%s' não encontrado em nenhum arquivo de %s", handlerName, handlerFolder)
	}

	return nil
}

func createDtoContentInput(dtoName string) string {
	return fmt.Sprintf(`	var input dto.%s
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}`, dtoName)
}

func createDtoContentOutput(dtoName string) string {
	return fmt.Sprintf(`	output := dto.%s{
		// preencha os campos aqui
	}
	c.JSON(200, output)`, dtoName)
}
