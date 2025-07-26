package utils

type Config struct {
	RoutesFile string `json:"routes_file"`

	HandlersFile   string `json:"handlers_file"`
	HandlersFolder string `json:"handlers_folder"`

	MiddlewaresFolder string `json:"middlewares_folder"`
	MiddlewaresFile   string `json:"middlewares_file"`

	DtoFolder string `json:"dto_folder"`
	DtoFile   string `json:"dto_file"`

	ServicesFolder   string `json:"services_folder"`
	ModelsFolder     string `json:"models_folder"`
	MigrationsFolder string `json:"migrations_folder"`
	DatabaseFolder   string `json:"database_folder"`

	Framework   string `json:"framework"`
	ProjectName string `json:"project_name"`
}
