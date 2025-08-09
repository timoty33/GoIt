package lint

import (
	"fmt"
	"goit/utils"
	"sync"
)

func RunStaticFmt(configProject utils.ConfigProject) error {
	if !isStaticCheckAvaible() {
		if err := installStaticCheck(); err != nil {
			return fmt.Errorf("erro ao instalar o static: %w", err)
		}
	}

	var wg sync.WaitGroup
	msgChan := make(chan error, 2)

	if configProject.Run.Lint.Format {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := goFmt(configProject.Run.Lint.LintBackEnd); err != nil {
				msgChan <- fmt.Errorf("erro ao formatar: %v", err)
			}
		}()
	}

	if configProject.Run.Lint.Lint {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := staticLint(configProject.Run.Lint.LintBackEnd); err != nil {
				msgChan <- fmt.Errorf("erro ao rodar linter: %v", err)
			}
		}()
	}

	wg.Wait()
	close(msgChan)

	for msg := range msgChan {
		if msg != nil {
			return msg
		}
	}

	return nil
}
