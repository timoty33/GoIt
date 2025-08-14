package dev

import (
	"fmt"
	"goit/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/google/shlex"

	"github.com/fsnotify/fsnotify"
)

func RunDevBackend(configProject utils.ConfigProject) error {
	fmt.Println("Iniciando BackEnd em modo de dev!")

	// hotreload
	if configProject.Run.Dev.HotReloadBackend.Active {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		listenPath := configProject.Run.Dev.HotReloadBackend.ListenPath

		// adiciona subpastas
		filepath.WalkDir(listenPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// ve se é um diretório
			if d.IsDir() {
				return watcher.Add(path)
			}

			return nil
		})

		go func() { // roda em paralelo para receber os valores das chans, e para fazer o loop (hotreload)
			for {
				select {
				// verificamos se existe um evento
				case event := <-watcher.Events:
					// verifica se é a criação de uma pasta
					if event.Op&fsnotify.Create == fsnotify.Create {
						info, err := os.Stat(event.Name)
						if err == nil && info.IsDir() {
							watcher.Add(event.Name)
						}
					}

					// valida o event, filtra ele
					if !shouldIgnore(event.Name, configProject.Run.Dev.Ignore) {
						command, err := shlex.Split(configProject.Run.Dev.InitCommandBackend)
						if err != nil {
							log.Panic(err)
						}
						err = utils.CmdExecute(command[0], command[1:]...)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}()
	}

	return nil
}
