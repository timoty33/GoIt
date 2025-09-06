package create

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"path/filepath"
)

func CreateModelNewFile(modelName, modelNameVerify string, configs utils.ConfigPaths, configsProject utils.ConfigProject, camps []string) error {
	if configsProject.Orm == "gorm" {
		content := fmt.Sprintf(`package model

import (
	"gorm.io/gorm"
)

%s
`, createModelStringGorm(modelNameVerify, camps))

		err := file.CreateArqVerify(configs.ModelsFolder, filepath.Join(configs.ModelsFolder, modelName+".go"), modelName, content)
		if err != nil {
			return fmt.Errorf("❌ Erro ao criar modelo: %w", err)
		}

	} else {
		content := fmt.Sprintf(`package model

%s
`, createModelString(modelNameVerify, camps))

		err := file.CreateArqVerify(configs.ModelsFolder, filepath.Join(configs.ModelsFolder, modelName+".go"), modelName, content)
		if err != nil {
			return fmt.Errorf("❌ Erro ao criar modelo: %w", err)
		}
	}

	return nil
}

func createModelStringGorm(name string, camps []string) string {
	content := "type " + name + " struct {\n\tgorm.Model\n"
	for _, param := range camps {
		content += "\t" + param + "\n"
	}
	content += "}\n"

	return content
}

func createModelString(name string, camps []string) string {
	content := "type " + name + " struct {\n"
	for _, camp := range camps {
		content += "\t" + camp + "\n"
	}
	content += "}\n"

	return content
}
