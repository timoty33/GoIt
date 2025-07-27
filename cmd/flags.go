package cmd

// init
var framework string
var install bool

// create
var newFile bool
var method string
var routeName string
var handlerName string
var dtoRoute string
var handlerParams []string

func init() {
	// goit init
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&framework, "framework", "f", "", "Framework a ser utilizado (ex: gin)")
	initCmd.Flags().BoolVarP(&install, "install", "i", false, "Instala as dependências após a criação do projeto")
	initCmd.MarkFlagRequired("framework")

	// goit create
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVar(&newFile, "new-file", false, "O código será criado em um novo arquivo ou em um arquivo existente?")

	createCmd.Flags().StringVar(&routeName, "route-name", "", "Modifica o nome da rota que será criada, originalmente é o mesmo nome do handler.")
	createCmd.Flags().StringVar(&handlerName, "handler-name", "", "Modifica o nome do handler que será atribuído à rota.")
	createCmd.Flags().StringSliceVar(&handlerParams, "params", []string{}, "Define parâmetros que serão automaticamente adicionados no handler.")
	createCmd.Flags().StringVar(&method, "method", "", "Define o método para a rota.")

	createCmd.Flags().StringVar(&dtoRoute, "dto-route", "", "Define qual rota usará aquele DTO.")

}
