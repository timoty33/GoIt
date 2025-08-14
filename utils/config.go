package utils

type LintType struct {
	Lint         bool   `json:"lint"`
	LintApply    bool   `json:"lint_apply"`
	Format       bool   `json:"format"`
	LintFrontEnd string `json:"frontend_path"` // caminho que será passado no biome
	LintBackEnd  string `json:"backend_path"`  // caminho que será passado no linter
}
type hotreloadBackend struct {
	Active     bool   `json:"active"`
	ListenPath string `json:"listen_path"`
}
type hotreloadFrontend struct {
	Active     bool   `json:"active"`
	ListenPath string `json:"listen_path"`
}
type Dev struct {
	HotReloadBackend   hotreloadBackend  `json:"hot_reload_backend"`
	HotReloadFrontend  hotreloadFrontend `json:"hot_reload_frontend"`
	InitCommandBackend string            `json:"init_command_backend"`
	Ignore             []string          `json:"ignore_paths"`
}
type Run struct {
	Dev  Dev      `json:"dev"`
	Lint LintType `json:"linter"`
}

type ConfigProject struct {
	Framework           string `json:"framework"`
	DataBase            string `json:"database"`
	Orm                 string `json:"orm"`
	Port                string `json:"port"`
	ProgrammingLanguage string `json:"programming_language"`
	ProjectName         string `json:"project_name"`
	ProjectType         string `json:"project_type"`

	Run Run `json:"run"`
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
