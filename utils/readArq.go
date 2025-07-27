package utils

import (
	"fmt"
	"os"
)

func ReadFile(fullPath string) (string, error) {
	contentBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("❌ Erro ao ler o arquivo '%s': %w", fullPath, err)
	}
	content := string(contentBytes)

	return content, err
}
