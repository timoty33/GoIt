package goitPluginOrganizer

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"os"
	"os/exec"
	"path/filepath"
)

// Insere o plugin no GPO com o import e o AddCommand
func loadPluginToGPO(namePlugin, fullpath string) error {
	contentFileGpo, err := file.ReadFile(fullpath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo GPO: %v", err)
	}

	// Insere o import
	contentAfterImport := "\t" + `"goit/plugins/` + namePlugin + `"` + "\n"
	newFileContent, err := utils.InsertAfterPlaceholder(contentFileGpo, "// goit:add-imports-here", contentAfterImport)
	if err != nil {
		return fmt.Errorf("erro ao inserir o import no GPO: %v", err)
	}

	// Insere o AddCommand
	contentAfterCommand := "\tgpo.AddCommand(" + namePlugin + ".Start" + ")\n"
	newFileContentWithPlugin, err := utils.InsertAfterPlaceholder(newFileContent, "// goit:add-plugins-here", contentAfterCommand)
	if err != nil {
		return fmt.Errorf("erro ao inserir o plugin no GPO: %v", err)
	}

	// Escreve o resultado final no arquivo
	err = os.WriteFile(fullpath, []byte(newFileContentWithPlugin), 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo GPO: %v", err)
	}

	return nil
}

// Mata o processo atual e recompila o GPO
func recompileGpo() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	mainGo := filepath.Join(exeDir, "gpo", "main.go")
	batFile := filepath.Join(exeDir, "recompile.bat")

	batContent := `
@echo off
timeout /t 1 > nul
echo 🛠️ Recompilando o GPO...
del "` + exePath + `"
go build -o "` + exePath + `" "` + mainGo + `"
if errorlevel 1 (
  echo ❌ Falha na compilação!
  pause
  exit /b 1
)
echo ✅ GPO recompilado com sucesso!
del "%~f0"
`

	err = os.WriteFile(batFile, []byte(batContent), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("cmd", "/C", "start", "", batFile)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	// Encerra o processo atual (GPO) para liberar o binário para recompilação
	os.Exit(0)
}
