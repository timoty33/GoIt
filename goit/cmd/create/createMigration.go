package create

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"path/filepath"
)

func CreateMigration(name, nameVerify, modelName string, configs utils.ConfigPaths, configsProject utils.ConfigProject) error {
	if configsProject.Orm == "gorm" {
		content := createMigrationFileContentGorm(nameVerify, modelName, configs, configsProject)

		err := file.CreateArqVerify(configs.MigrationsFolder, filepath.Join(configs.MigrationsFolder, name+".go"), name+".go", content)
		if err != nil {
			return fmt.Errorf("❌ Erro ao criar migration: %w", err)
		}
	} else {
		content := createMigrationFileContent(nameVerify)

		err := file.CreateArqVerify(configs.MigrationsFolder, filepath.Join(configs.MigrationsFolder, name+".go"), name+".go", content)
		if err != nil {
			return fmt.Errorf("❌ Erro ao criar migration: %w", err)
		}
	}

	return nil
}

func createMigrationFileContentGorm(nameVerify, modelName string, configs utils.ConfigPaths, configsProject utils.ConfigProject) string {
	return `package migrations

import (
    "gorm.io/gorm"

	"` + configsProject.ProjectName + "/" + configs.ModelsFolder + `"
)

func ` + nameVerify + `(db *gorm.DB) error {
    return db.Migrator().CreateTable(&model.` + modelName + `{})
}
`
}

func createMigrationFileContent(nameVerify string) string {
	return `package migrations

import (
    "database/sql"
)

func ` + nameVerify + `(db *sql.DB) error {
    query := 

    _, err := db.Exec(query)
    return err
}
`
}
