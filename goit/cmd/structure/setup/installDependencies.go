package setup

import (
	"fmt"
	"goit/utils"
)

func InstallDependenciesGo(projectPath, framework, db string) error {
	switch framework {
	case "Gin":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "github.com/gin-gonic/gin")
		if err != nil {
			return fmt.Errorf("erro ao instalar Gin: %w", err)
		}

	case "Echo":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "github.com/labstack/echo/v4")
		if err != nil {
			return fmt.Errorf("erro ao instalar Echo: %w", err)
		}

	case "Fiber":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "github.com/gofiber/fiber/v2")
		if err != nil {
			return fmt.Errorf("erro ao instalar Fiber: %w", err)
		}
	}

	switch db {
	case "PostgreSQL":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "gorm.io/driver/postgres", "gorm.io/gorm")
		if err != nil {
			return fmt.Errorf("erro ao instalar Gorm com PostgreSQL: %w", err)
		}

	case "MySQL":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "gorm.io/driver/mysql", "gorm.io/gorm")
		if err != nil {
			return fmt.Errorf("erro ao instalar Gorm com MySQL: %w", err)
		}

	case "SQLite":
		err := utils.CmdExecuteInDir(projectPath, "go", "get", "gorm.io/driver/sqlite", "gorm.io/gorm")
		if err != nil {
			return fmt.Errorf("erro ao instalar Gorm com SQLite: %w", err)
		}

		// case "MongoDB":
		// 	err := installGoMongo()
		// 	if err != nil {
		// 		return fmt.Errorf("erro ao instalar Gorm com MongoDB: %w", err)
		// 	}

	}

	return nil
}

func InstallDependenciesJS(projectPath, framework, db string) error {
	// Aqui você pode adicionar a lógica para instalar dependências do JavaScript
	// Exemplo: npm install express mongoose
	return nil
}

func InstallDependenciesTS(projectPath, framework, db string) error {
	// Aqui você pode adicionar a lógica para instalar dependências do TypeScript
	// Exemplo: npm install typescript @types/node
	return nil
}
