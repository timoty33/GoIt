package lint

import (
	"fmt"
	"goit/utils"
	"goit/utils/file"
)

func RunBiome(configProject utils.ConfigProject) error {
	biomeInstall := isBiomeAvaible()
	if !biomeInstall {
		if err := installBiomeNpm(); err != nil {
			return fmt.Errorf("erro ao instalar o biome: %w", err)
		}
		if err := biomeInit(); err != nil {
			return fmt.Errorf("erro ao inicializar biome: %w", err)
		}

		if configProject.Run.Lint.LintApply {
			if err := biomeRunApply(configProject); err != nil {
				return fmt.Errorf("erro ao usar o biome: %w", err)
			}

			return nil
		}
		if err := biomeRun(configProject); err != nil {
			return fmt.Errorf("erro ao usar o biome: %w", err)
		}

		return nil
	}

	if !file.FileExists("biome.json") {
		fmt.Println("⚙️ biome.json não encontrado, inicializando...")
		if err := biomeInit(); err != nil {
			return fmt.Errorf("erro ao inicializar biome: %w", err)
		}
	}

	if configProject.Run.Lint.LintApply {
		if err := biomeRunApply(configProject); err != nil {
			return fmt.Errorf("erro ao usar o biome: %w", err)
		}

		return nil
	}
	if err := biomeRun(configProject); err != nil {
		return fmt.Errorf("erro ao usar o biome: %w", err)
	}

	return nil
}
