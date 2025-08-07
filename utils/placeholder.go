package utils

import (
	"fmt"
	"strings"
)

// InsertAfterPlaceholder encontra um marcador de posição e insere um novo conteúdo após ele.
func InsertAfterPlaceholder(fileContent, placeholder, contentToAdd string) (string, error) {

	if !strings.Contains(fileContent, placeholder) {
		return "", fmt.Errorf("marcador '%s' não encontrado", placeholder)
	}

	// Prepara a string de substituição: o próprio placeholder, uma nova linha e o conteúdo.
	replacement := placeholder + "\n" + contentToAdd

	// Se existe, realiza a substituição com segurança.
	return strings.Replace(fileContent, placeholder, replacement, 1), nil
}
