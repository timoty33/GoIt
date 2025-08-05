package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigProjectFileName = ".goit.config.json"
	ConfigPathsFileName   = ".goit.config.paths.json"
)

func SaveJsonConfigProject(configProject ConfigProject, projectPath string) error {
	fullPathProject := filepath.Join(projectPath, ConfigProjectFileName)

	arquivoConfigProject, err := os.Create(fullPathProject)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo de configuração do projeto '%s': %w", fullPathProject, err)
	}
	defer arquivoConfigProject.Close()

	encoder := json.NewEncoder(arquivoConfigProject)
	encoder.SetIndent("", "  ") // Deixa o JSON formatado para melhor leitura
	if err := encoder.Encode(configProject); err != nil {
		return fmt.Errorf("erro ao escrever JSON no arquivo de configuração do projeto: %w", err)
	}

	return nil
}

func SaveJsonConfigPaths(configPaths ConfigPaths, projectPath string) error {
	fullPathPaths := filepath.Join(projectPath, ConfigPathsFileName)

	arquivoConfigPaths, err := os.Create(fullPathPaths)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo de configuração dos caminhos '%s': %w", fullPathPaths, err)
	}
	defer arquivoConfigPaths.Close()

	encoder := json.NewEncoder(arquivoConfigPaths)
	encoder.SetIndent("", "  ") // Deixa o JSON formatado para melhor leitura
	if err := encoder.Encode(configPaths); err != nil {
		return fmt.Errorf("erro ao escrever JSON no arquivo de configuração dos caminhos: %w", err)
	}

	return nil
}

func SaveJsonConfigs(configProject ConfigProject, configPaths ConfigPaths, projectPath string) error {

	err := SaveJsonConfigPaths(configPaths, projectPath)
	if err != nil {
		return err
	}

	err = SaveJsonConfigProject(configProject, projectPath)
	if err != nil {
		return err
	}

	return nil
}

// LoadJsonConfig carrega a configuração do projeto a partir do diretório atual.
func LoadJsonConfig() (ConfigProject, ConfigPaths, error) {
	var configProject ConfigProject
	var configPaths ConfigPaths

	arquivoConfigProject, err := os.Open(ConfigProjectFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return configProject, configPaths, fmt.Errorf("arquivo de configuração do projeto '%s' não encontrado", ConfigProjectFileName)
		}
		return configProject, configPaths, fmt.Errorf("erro ao abrir o arquivo de configuração do projeto: %w", err)
	}
	defer arquivoConfigProject.Close()

	arquivoConfigPaths, err := os.Open(ConfigPathsFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return configProject, configPaths, fmt.Errorf("arquivo de configuração dos caminhos '%s' não encontrado", ConfigPathsFileName)
		}
		return configProject, configPaths, fmt.Errorf("erro ao abrir o arquivo de configuração do projeto: %w", err)
	}
	defer arquivoConfigPaths.Close()

	decoder := json.NewDecoder(arquivoConfigProject)
	if err := decoder.Decode(&configProject); err != nil {
		return configProject, configPaths, fmt.Errorf("erro ao decodificar o arquivo de configuração do projeto: %w", err)
	}

	decoder = json.NewDecoder(arquivoConfigPaths)
	if err := decoder.Decode(&configPaths); err != nil {
		return configProject, configPaths, fmt.Errorf("erro ao decodificar o arquivo de configuração dos caminhos: %w", err)
	}

	return configProject, configPaths, nil
}

// ===/===/===/===

func LoadJsonListString(fullpath string) ([]string, error) {
	var result []string

	file, err := os.Open(fullpath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar o json: %w", err)
	}

	return result, nil
}

func SaveJsonListString(list []string, fullpath string) error {
	file, err := os.Open(fullpath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(list); err != nil {
		return fmt.Errorf("erro ao codificar o json: %w", err)
	}
	return nil
}
