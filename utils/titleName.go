package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func TitleNameVerify(name string) (string, error) {
	var titleName string

	if !isValidIdentifier(name) {
		return titleName, fmt.Errorf("❌ Nome de handler inválido: '%s'. Use apenas letras e números, começando com uma letra", name)
	}

	// Capitaliza a primeira letra corretamente
	titleName = strings.Title(name)

	return titleName, nil
}

// isValidIdentifier verifica se o nome contém apenas letras ou números
// e começa com uma letra (padrão de identificadores em Go).
func isValidIdentifier(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		if i == 0 && !unicode.IsLetter(r) {
			return false
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
