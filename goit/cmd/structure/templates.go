package structure

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	ProjectName string
}

func RenderTemplates(templates map[string]string, data TemplateData, outputDir string) error {
	for path, content := range templates {
		// Parse o conteúdo do template
		tmpl, err := template.New(filepath.Base(path)).Parse(content)
		if err != nil {
			return err
		}

		// Remove a parte "../templates/" do caminho
		relPath := strings.TrimPrefix(path, "../templates/")
		// Remove a extensão .tmpl
		destPath := strings.TrimSuffix(relPath, ".tmpl")

		// Cria o caminho final do arquivo
		finalPath := filepath.Join(outputDir, destPath)

		// Cria diretórios, se não existirem
		err = os.MkdirAll(filepath.Dir(finalPath), os.ModePerm)
		if err != nil {
			return err
		}

		// Cria e escreve no arquivo final
		outFile, err := os.Create(finalPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		err = tmpl.Execute(outFile, data)
		if err != nil {
			return err
		}
	}

	return nil
}
