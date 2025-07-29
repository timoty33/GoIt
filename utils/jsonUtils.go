package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	permPasta   = 0755
	permArquivo = 0644
	// ConfigFileName é o nome padrão para o arquivo de configuração do GoIt.
	ConfigFileName = ".goit.config.json"
)

// SaveJsonConfigProject salva a configuração do projeto em um arquivo JSON na raiz do projeto.
func SaveJsonConfigProject(config ConfigProject, projectPath string) error {
	fullPath := filepath.Join(projectPath, ConfigFileName)

	arquivo, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo de configuração '%s': %w", fullPath, err)
	}
	defer arquivo.Close()

	encoder := json.NewEncoder(arquivo)
	encoder.SetIndent("", "  ") // Deixa o JSON formatado para melhor leitura
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("erro ao escrever JSON no arquivo de configuração: %w", err)
	}

	return nil
}

// LoadJsonConfig carrega a configuração do projeto a partir do diretório atual.
func LoadJsonConfig() (Config, error) {
	var config Config

	arquivo, err := os.Open(ConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return config, fmt.Errorf("arquivo de configuração '%s' não encontrado. Execute este comando na raiz do seu projeto GoIt", ConfigFileName)
		}
		return config, fmt.Errorf("erro ao abrir o arquivo de configuração: %w", err)
	}
	defer arquivo.Close()

	decoder := json.NewDecoder(arquivo)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("erro ao decodificar o arquivo de configuração JSON: %w", err)
	}

	return config, nil
}
