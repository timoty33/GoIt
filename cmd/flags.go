package cmd

// create
var newFile bool
var method string
var routeName string
var handlerName string
var handlerParams []string
var camps []string
var dtoMode string
var modelName string

func init() {
	// goit init
	rootCmd.AddCommand(initCmd)

	// goit create
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&newFile, "new-file", "n", false, "O código será criado em um novo arquivo ou em um arquivo existente?")

	// handler/route
	createCmd.Flags().StringVarP(&routeName, "route-name", "R", "", "Modifica o nome da rota que será criada, originalmente é o mesmo nome do handler.")
	createCmd.Flags().StringVarP(&handlerName, "for", "H", "", "Modifica o nome do handler que será atribuído à rota.") // Também é usado pelo dto para saber o handler
	createCmd.Flags().StringSliceVarP(&handlerParams, "params", "p", []string{}, "Define parâmetros que serão automaticamente adicionados no handler.")
	createCmd.Flags().StringVarP(&method, "method", "M", "", "Define o método para a rota.")

	// dto
	createCmd.Flags().StringSliceVarP(&camps, "camps", "c", []string{}, "Define os nomes dos campos da struct.") // dto/model
	createCmd.Flags().StringVarP(&dtoMode, "dto-mode", "m", "input", "Define o modo que o DTO será usado: input/output")

	// migration
	createCmd.Flags().StringVar(&modelName, "model", "", "Define o nome do modelo que será usado na migration.")
}
