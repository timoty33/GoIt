package dev

import (
	"goit/utils"
)

func RunDevFullstack(config utils.ConfigProject) error {
	errCh := make(chan error, 2)

	go func() {
		errCh <- RunDevFrontend(config)
	}()

	go func() {
		errCh <- RunDevBackend(config)
	}()

	// Espera ambos terminarem ou retornarem erro
	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}
