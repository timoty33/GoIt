package lint

import (
	"fmt"
	"goit/utils"
	"sync"
)

func RunStaticFmt(configProject utils.ConfigProject) error {
	if !IsStaticCheckAvaible() {
		if err := InstallStaticCheck(); err != nil {
			return fmt.Errorf("erro ao instalar o static: %w", err)
		}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	if configProject.Run.Lint.Format {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := GoFmt(configProject.Run.Lint.LintBackEnd); err != nil {
				errChan <- fmt.Errorf("erro ao formatar: %w", err)
			}
		}()
	}

	if configProject.Run.Lint.Lint {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := StaticLint(configProject.Run.Lint.LintBackEnd); err != nil {
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
