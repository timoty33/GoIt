package file

import (
	"os"
	"path/filepath"
	"strings"
)

// PercorrerDiretorio lê todos os arquivos .tmpl da pasta e retorna um map[path]content
func PercorrerDiretorio(dir string) (map[string]string, error) {
	templates := make(map[string]string)

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(d.Name(), ".tmpl") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Salva no map: chave = caminho relativo a partir da pasta templates
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}

			templates[relPath] = string(content)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return templates, nil
}
