package config

import (
	"fmt"
	"goit/utils"
)

func UpdateConfigKey(configFile bool, key, value string) error {
	// Carrega as configurações existentes
	configProj, configPaths, err := utils.LoadJsonConfig()
	if err != nil {
		return err
	}

	if configFile {
		// Atualiza com base na chave
		switch key {
		case "framework":
			configProj.Framework = value
		case "database":
			configProj.DataBase = value
		case "orm":
			configProj.Orm = value
		case "port":
			configProj.Port = value
		case "language":
			configProj.ProgrammingLanguage = value
		case "name":
			configProj.ProjectName = value
		default:
			return fmt.Errorf("chave de configuração desconhecida: %s", key)
		}
	} else {
		switch key {
		case "routesFile":
			configPaths.RoutesFile = value
		case "handlersFile":
			configPaths.HandlersFile = value
		case "handlersFolder":
			configPaths.HandlersFolder = value
		case "middlewaresFolder":
			configPaths.MiddlewaresFolder = value
		case "middlewareFile":
			configPaths.MiddlewaresFile = value
		case "dtoFolder":
			configPaths.DtoFolder = value
		case "dtoFile":
			configPaths.DtoFile = value
		case "modelsFolder":
			configPaths.ModelsFolder = value
		case "servicesFolder":
			configPaths.ServicesFolder = value
		case "migrationsFolder":
			configPaths.MigrationsFolder = value
		case "repositoryFolder":
			configPaths.RepositoryFolder = value
		case "databaseFolder":
			configPaths.DatabaseFolder = value
		default:
			return fmt.Errorf("chave de configuração desconhecida: %s", key)
		}
	}

	// Salva novamente as configs
	projectPath := "." // ou caminho específico
	return utils.SaveJsonConfigs(configProj, configPaths, projectPath)
}
