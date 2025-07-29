package structure

import "goit/utils"

func CreateGoitConfigs(projectName, framework, ProgrammingLanguage, database, orm string) error {
	projectConfig := utils.ConfigProject{
		Framework:           framework,
		DataBase:            database,
		Orm:                 orm,
		ProgrammingLanguage: ProgrammingLanguage,
		ProjectName:         projectName,
	}
}