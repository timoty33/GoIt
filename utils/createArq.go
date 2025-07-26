package utils

import (
	"fmt"
	"os"
)

func CreateArqVerify(folder, fullPath, nome, content string) error {
	// Verifica se diretório existe, se não, cria
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", folder, err)
	}

	// Verifica duplicação
	if _, err := os.Stat(fullPath); err == nil {
		return fmt.Errorf("arquivo '%s' já existe", nome)
	}

	// Cria o arquivo e escreve o content
	if err := os.WriteFile(fullPath, []byte(content), permArquivo); err != nil {
		return fmt.Errorf("erro ao criar arquivo %s: %w", fullPath, err)
	}

	return nil
}
