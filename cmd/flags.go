package cmd

// init
var framework string
var install bool

// create
var newFile bool
var method string
var routeName string
var dtoRoute string

func init() {
	// goit init
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&framework, "framework", "f", "", "Framework a ser utilizado (ex: gin)")
	initCmd.Flags().BoolVarP(&install, "install", "i", false, "Instala as dependências após a criação do projeto")
	initCmd.MarkFlagRequired("framework")

	// goit create
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&newFile, "new-file", "n", false, "O código será criado em um novo arquivo ou em um arquivo existente?")
	createCmd.Flags().StringVarP(&routeName, "route-name", "r", "", "Modifica o nome da rota que será criada, originalmente é o mesmo nome do handler.")
	createCmd.Flags().StringVarP(&method, "method", "m", "", "Define o método para a rota, ou se estiver criando o handler, a rota será automaticamente criada com o método recebido.")
	createCmd.Flags().StringVarP(&dtoRoute, "dto-route", "d", "", "Define qual rota usará aquele DTO.")
}
