package cmd

// create
var (
	newFile       bool
	method        string
	routeName     string
	handlerName   string
	handlerParams []string
	camps         []string
	dtoMode       string
	modelName     string
)

// run
var (
	runOnlyBackend  bool
	runOnlyFrontend bool
)

func init() {
	// goit init
	rootCmd.AddCommand(initCmd)

	// goit create
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&newFile, "new-file", "n", false, "O código será criado em um novo arquivo ou em um arquivo existente?")

	// handler
	createCmd.AddCommand(createHandler, createRoute, createEnv, createMiddleware)

	createHandler.Flags().StringVarP(&routeName, "route-name", "R", "", "Modifica o nome da rota que será criada, originalmente é o mesmo nome do handler.")
	createHandler.Flags().StringSliceVarP(&handlerParams, "params", "p", []string{}, "Define parâmetros que serão automaticamente adicionados no handler.")
	createHandler.Flags().StringVarP(&method, "method", "M", "", "Define o método para a rota.")

	// route
	createRoute.Flags().StringVarP(&handlerName, "for", "H", "handlerName", "Modifica o nome do handler que será atribuído à rota.") // Também é usado pelo dto para saber o handler
	createRoute.Flags().StringVarP(&method, "method", "M", "get", "Define o método para a rota.")

	// dto
	createCmd.Flags().StringSliceVarP(&camps, "camps", "c", []string{}, "Define os nomes dos campos da struct.") // dto/model
	createCmd.Flags().StringVarP(&dtoMode, "dto-mode", "m", "input", "Define o modo que o DTO será usado: input/output")

	// migration
	createCmd.Flags().StringVar(&modelName, "model", "", "Define o nome do modelo que será usado na migration.")

	// run
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&runOnlyBackend, "backend", "b", false, "Roda apenas o backend.")
	runCmd.Flags().BoolVarP(&runOnlyFrontend, "frontend", "f", false, "Roda apenas o frontend.")
}
