package lint

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
	"sync"
)

func RunBiome(configProject utils.ConfigProject) error {
	biomeInstall := IsBiomeAvaible()
	if !biomeInstall {
		if err := InstallBiomeNpm(); err != nil {
			return fmt.Errorf("erro ao instalar o biome: %w", err)
		}
		if err := BiomeInit(); err != nil {
			return fmt.Errorf("erro ao inicializar biome: %w", err)
		}

		// Rodar formatter e linter
		if configProject.Run.Lint.LintApply {
			// NÃO roda em paralelo, pois os dois modificam arquivos
			if configProject.Run.Lint.Format {
				if err := BiomeFormatter(configProject.Run.Lint.LintFrontEnd); err != nil {
					return fmt.Errorf("erro ao formatar: %w", err)
				}
			}
			if err := BiomeLintApply(configProject.Run.Lint.LintFrontEnd); err != nil {
				return fmt.Errorf("erro ao aplicar o linter: %w", err)
			}
			return nil
		}

		var wg sync.WaitGroup
		errChan := make(chan error, 2)

		if configProject.Run.Lint.Format {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := BiomeFormatter(configProject.Run.Lint.LintFrontEnd); err != nil {
					errChan <- fmt.Errorf("erro ao formatar: %w", err)
				}
			}()
		}

		if configProject.Run.Lint.Lint {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := BiomeLint(configProject.Run.Lint.LintFrontEnd); err != nil {
					errChan <- fmt.Errorf("erro ao rodar linter: %w", err)
				}
			}()
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			if err != nil {
				return err
			}
		}

		return nil
	}

	if !file.FileExists("biome.json") {
		fmt.Println("⚙️ biome.json não encontrado, inicializando...")
		if err := BiomeInit(); err != nil {
			return fmt.Errorf("erro ao inicializar biome: %w", err)
		}
	}

	// Rodar formatter e linter
	if configProject.Run.Lint.LintApply {
		// NÃO roda em paralelo, pois os dois modificam arquivos
		if configProject.Run.Lint.Format {
			if err := BiomeFormatter(configProject.Run.Lint.LintFrontEnd); err != nil {
				return fmt.Errorf("erro ao formatar: %w", err)
			}
		}
		if err := BiomeLintApply(configProject.Run.Lint.LintFrontEnd); err != nil {
			return fmt.Errorf("erro ao aplicar o linter: %w", err)
		}
		return nil
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	if configProject.Run.Lint.Format {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := BiomeFormatter(configProject.Run.Lint.LintFrontEnd); err != nil {
				errChan <- fmt.Errorf("erro ao formatar: %w", err)
			}
		}()
	}

	if configProject.Run.Lint.Lint {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := BiomeLint(configProject.Run.Lint.LintFrontEnd); err != nil {
				errChan <- fmt.Errorf("erro ao rodar linter: %w", err)
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
