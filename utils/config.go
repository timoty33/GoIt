package utils

type ConfigProject struct {
	Framework           string `json:"framework"`
	DataBase            string `json:"database"`
	Orm                 string `json:"orm"`
	Port                string `json:"port"`
	ProgrammingLanguage string `json:"programming_language"`
	ProjectName         string `json:"project_name"`

	HotReload bool `json:"hot_reload"`
}

type ConfigPaths struct {
	ServerFile string `json:"server_file"`

	RoutesFile string `json:"routes_file"`

	HandlersFile   string `json:"handlers_file"`
	HandlersFolder string `json:"handlers_folder"`

	MiddlewaresFolder string `json:"middlewares_folder"`
	MiddlewaresFile   string `json:"middlewares_file"`

	DtoFolder string `json:"dto_folder"`
	DtoFile   string `json:"dto_file"`

	ModelsFolder string `json:"models_folder"`

	ServicesFolder string `json:"services_folder"`

	MigrationsFolder string `json:"migrations_folder"`

	RepositoryFolder string `json:"repository_folder"`
	DatabaseFolder   string `json:"database_folder"`
}

type Config struct {
	Project ConfigProject `json:"project"`
	Paths   ConfigPaths   `json:"paths"`
}
